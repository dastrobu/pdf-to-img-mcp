package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/scopweb/mcp-go-pdf-to-img/internal/config"
	"github.com/scopweb/mcp-go-pdf-to-img/internal/logging"
	"github.com/scopweb/mcp-go-pdf-to-img/pkg/converter"
)

// Converter interface for PDF operations
type Converter interface {
	Convert(opts *converter.ConvertOptions) (*converter.ConvertResult, error)
	GetPDFInfo(pdfPath string) (map[string]interface{}, error)
	Close() error
}

// Handlers encapsulates all the HTTP handlers and their dependencies.
type Handlers struct {
	converter Converter
	logger    logging.Logger
	config    *config.ServerConfig
}

// NewHandlers creates a new set of handlers.
func NewHandlers(conv Converter, logger logging.Logger, config *config.ServerConfig) *Handlers {
	return &Handlers{
		converter: conv,
		logger:    logger,
		config:    config,
	}
}

// Health responds with an OK status.
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// ConvertRequest holds the conversion parameters.
type ConvertRequest struct {
	DPI          int    `json:"dpi,omitempty"`
	Format       string `json:"format,omitempty"`
	StartPage    int    `json:"start_page,omitempty"`
	EndPage      int    `json:"end_page,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	RetryFailed  bool   `json:"retry_failed,omitempty"`
	PoolSize     int    `json:"pool_size,omitempty"`
	RefreshEvery int    `json:"refresh_every,omitempty"`
}

// ConvertResponse holds the result of the conversion.
type ConvertResponse struct {
	TotalPages   int      `json:"total_pages"`
	Successful   int      `json:"successful"`
	Failed       int      `json:"failed"`
	OutputFiles  []string `json:"output_files,omitempty"`
	Errors       []string `json:"errors,omitempty"`
	WarningPages []int    `json:"warning_pages,omitempty"`
}

// Convert converts a PDF to images and returns a ZIP with the results.
func (h *Handlers) Convert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(h.config.MaxUploadSize); err != nil {
		h.logger.Error("failed to parse multipart form", err)
		http.Error(w, "invalid request format", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		h.logger.Warn("missing file field", slog.Any("error", err))
		http.Error(w, "missing file field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file for the upload
	tmpFile, err := os.CreateTemp("", "upload-*.pdf")
	if err != nil {
		h.logger.Error("failed to create temp file", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	// Save the file
	if _, err := io.Copy(tmpFile, file); err != nil {
		tmpFile.Close()
		h.logger.Error("failed to save uploaded file", err)
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	tmpFile.Close()

	// Parse conversion parameters
	var req ConvertRequest
	if err := h.parseConvertParams(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a temporary output directory
	tmpOutputDir, err := os.MkdirTemp("", "pdf2img-*")
	if err != nil {
		h.logger.Error("failed to create temp output dir", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpOutputDir)

	// Configure conversion options
	opts := &converter.ConvertOptions{
		InputPath:    tmpPath,
		OutputDir:    tmpOutputDir,
		Format:       req.Format,
		DPI:          float64(req.DPI),
		StartPage:    req.StartPage,
		EndPage:      req.EndPage,
		Prefix:       req.Prefix,
		RetryFailed:  req.RetryFailed,
		MaxPoolSize:  req.PoolSize,
		RefreshEvery: req.RefreshEvery,
	}

	// Convert the PDF
	result, err := h.converter.Convert(opts)
	if err != nil {
		h.logger.Error("PDF conversion failed", err)
		http.Error(w, "failed to convert PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the response
	response := ConvertResponse{
		TotalPages:   result.TotalPages,
		Successful:   result.Successful,
		Failed:       result.Failed,
		OutputFiles:  result.OutputFiles,
		Errors:       result.Errors,
		WarningPages: result.WarningPages,
	}

	// If there are no output files, return the result as JSON
	if len(result.OutputFiles) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return a ZIP with the images
	h.serveZip(w, header.Filename, result.OutputFiles)
}

func (h *Handlers) parseConvertParams(r *http.Request, req *ConvertRequest) error {
	// Defaults
	req.Format = h.config.Convert.DefaultFormat
	if req.Format == "" {
		req.Format = "png"
	}
	req.DPI = h.config.Convert.DefaultDPI
	if req.DPI <= 0 {
		req.DPI = 150
	}
	req.PoolSize = h.config.Convert.DefaultPoolSize
	req.RefreshEvery = h.config.Convert.RefreshEvery
	req.RetryFailed = h.config.Convert.RetryOnFailure

	// Parse DPI
	if dpiStr := r.FormValue("dpi"); dpiStr != "" {
		if dpi, err := strconv.Atoi(dpiStr); err == nil && dpi > 0 {
			req.DPI = dpi
		}
	}

	// Parse the format
	if format := r.FormValue("format"); format != "" {
		req.Format = format
	}

	// Parse the pages
	if startStr := r.FormValue("start_page"); startStr != "" {
		if start, err := strconv.Atoi(startStr); err == nil {
			req.StartPage = start
		}
	}
	if endStr := r.FormValue("end_page"); endStr != "" {
		if end, err := strconv.Atoi(endStr); err == nil {
			req.EndPage = end
		}
	}

	// Parse the prefix
	req.Prefix = r.FormValue("prefix")
	if req.Prefix == "" {
		req.Prefix = "page_"
	}

	// Parse retry
	if retryStr := r.FormValue("retry"); retryStr != "" {
		req.RetryFailed, _ = strconv.ParseBool(retryStr)
	}

	// Parse pool size
	if poolStr := r.FormValue("pool_size"); poolStr != "" {
		if pool, err := strconv.Atoi(poolStr); err == nil && pool > 0 {
			req.PoolSize = pool
		}
	}

	// Parse refresh every
	if refreshStr := r.FormValue("refresh_every"); refreshStr != "" {
		if refresh, err := strconv.Atoi(refreshStr); err == nil && refresh >= 0 {
			req.RefreshEvery = refresh
		}
	}

	return nil
}

func (h *Handlers) serveZip(w http.ResponseWriter, filename string, files []string) {
	// Create a temporary ZIP file
	zipFile, err := os.CreateTemp("", "result-*.zip")
	if err != nil {
		h.logger.Error("failed to create zip file", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	zipPath := zipFile.Name()
	zipFile.Close()
	defer os.Remove(zipPath)

	// Create the ZIP
	if err := h.createZip(zipPath, files); err != nil {
		h.logger.Error("failed to create zip", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Open the ZIP file
	zipRead, err := os.Open(zipPath)
	if err != nil {
		h.logger.Error("failed to open zip", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer zipRead.Close()

	// Send the response
	w.Header().Set("Content-Type", "application/zip")
	zipName := sanitizeFilename(filename) + "-images.zip"
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))

	if _, err := io.Copy(w, zipRead); err != nil {
		h.logger.Error("error writing zip response", err)
	}
}

func (h *Handlers) createZip(zipPath string, files []string) error {
	f, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	for _, filePath := range files {
		fw, err := zw.Create(filepath.Base(filePath))
		if err != nil {
			continue
		}
		fr, err := os.Open(filePath)
		if err != nil {
			continue
		}
		io.Copy(fw, fr)
		fr.Close()
	}
	return nil
}

// Info returns information about a PDF.
func (h *Handlers) Info(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tmpPath string

	if r.Method == http.MethodPost {
		// Parse multipart form
		if err := r.ParseMultipartForm(h.config.MaxUploadSize); err != nil {
			h.logger.Error("failed to parse multipart form", err)
			http.Error(w, "invalid request format", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			h.logger.Warn("missing file field", slog.Any("error", err))
			http.Error(w, "missing file field", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a temporary file for the upload
		tmpFile, err := os.CreateTemp("", "upload-*.pdf")
		if err != nil {
			h.logger.Error("failed to create temp file", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		tmpPath = tmpFile.Name()
		defer os.Remove(tmpPath)

		if _, err := io.Copy(tmpFile, file); err != nil {
			tmpFile.Close()
			h.logger.Error("failed to save uploaded file", err)
			http.Error(w, "failed to save file", http.StatusInternalServerError)
			return
		}
		tmpFile.Close()
	} else {
		// GET: obtain the file URL as a parameter
		fileURL := r.FormValue("url")
		if fileURL == "" {
			http.Error(w, "missing url parameter", http.StatusBadRequest)
			return
		}

		// Download the file
		tmpFile, err := os.CreateTemp("", "upload-*.pdf")
		if err != nil {
			h.logger.Error("failed to create temp file", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		tmpPath = tmpFile.Name()
		defer os.Remove(tmpPath)

		resp, err := http.Get(fileURL)
		if err != nil {
			tmpFile.Close()
			h.logger.Error("failed to download file", err)
			http.Error(w, "failed to download file", http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		if _, err := io.Copy(tmpFile, resp.Body); err != nil {
			tmpFile.Close()
			h.logger.Error("failed to save downloaded file", err)
			http.Error(w, "failed to save file", http.StatusInternalServerError)
			return
		}
		tmpFile.Close()
	}

	// Get the PDF information
	info, err := h.converter.GetPDFInfo(tmpPath)
	if err != nil {
		h.logger.Error("failed to get PDF info", err)
		http.Error(w, "failed to get PDF info: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// sanitizeFilename cleans a filename to avoid problematic characters.
func sanitizeFilename(filename string) string {
	if ext := filepath.Ext(filename); ext != "" {
		filename = filename[:len(filename)-len(ext)]
	}
	return filepath.Base(filename)
}
