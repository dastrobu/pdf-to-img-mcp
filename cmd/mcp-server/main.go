package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	localmcp "github.com/scopweb/mcp-go-pdf-to-img/mcp"
)

const (
	serverName    = "pdf2img"
	serverVersion = "1.0.0"
)

func main() {
	log.Println("🚀 Starting PDF2IMG MCP Server")

	// Create the local MCP implementation (holds the converter/splitter/compress
	// logic) so the go-sdk server can delegate the actual work to it.
	localServer, err := localmcp.NewMCPServer()
	if err != nil {
		log.Fatalf("Failed to create local MCP server: %v", err)
	}
	defer localServer.Close()

	// Create the official go-sdk MCP server.
	srv := mcp.NewServer(
		&mcp.Implementation{
			Name:    serverName,
			Version: serverVersion,
		},
		nil,
	)

	// Register tools.
	registerTools(srv, localServer)

	log.Println("✅ Server ready - Waiting for connections...")

	// Serve over stdio.
	if err := srv.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// PdfToImagesInput defines the input parameters for the pdf_to_images tool.
type PdfToImagesInput struct {
	PDFPath   string  `json:"pdf_path" jsonschema:"Path to the PDF file to convert"`
	OutputDir string  `json:"output_dir" jsonschema:"Directory where images will be saved"`
	Format    string  `json:"format,omitempty" jsonschema:"Output format: 'png' or 'jpg' (default: png)"`
	DPI       float64 `json:"dpi,omitempty" jsonschema:"DPI for rendering (default: 150)"`
	StartPage int     `json:"start_page,omitempty" jsonschema:"Start page number (1-indexed, 0 for first page)"`
	EndPage   int     `json:"end_page,omitempty" jsonschema:"End page number (1-indexed, 0 for last page)"`
	Prefix    string  `json:"prefix,omitempty" jsonschema:"Prefix for output filenames (default: page_)"`
}

// PdfInfoInput defines the input parameters for the pdf_info tool.
type PdfInfoInput struct {
	PDFPath string `json:"pdf_path" jsonschema:"Path to the PDF file"`
}

// registerTools registers all PDF2IMG tools with the MCP server, delegating the
// actual work to the local (dependency-free) MCP implementation.
func registerTools(srv *mcp.Server, localServer *localmcp.MCPServer) {
	// pdf_to_images
	mcp.AddTool(srv,
		&mcp.Tool{
			Name:        "pdf_to_images",
			Description: "Convert PDF pages to PNG or JPG images",
			InputSchema: GenerateSchema[PdfToImagesInput](),
			Annotations: &mcp.ToolAnnotations{
				Title:           "PDF to Images",
				ReadOnlyHint:    false,
				IdempotentHint:  true,
				DestructiveHint: new(false),
				OpenWorldHint:   new(true),
			},
		},
		func(ctx context.Context, request *mcp.CallToolRequest, input PdfToImagesInput) (*mcp.CallToolResult, any, error) {
			args, err := json.Marshal(input)
			if err != nil {
				return nil, nil, fmt.Errorf("marshaling input: %w", err)
			}
			result, err := localServer.ExecuteTool("pdf_to_images", args)
			if err != nil {
				return nil, nil, fmt.Errorf("conversion error: %w", err)
			}
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: result.Content}},
			}, nil, nil
		},
	)

	// pdf_info
	mcp.AddTool(srv,
		&mcp.Tool{
			Name:        "pdf_info",
			Description: "Get information about a PDF file (pages, size, dimensions)",
			InputSchema: GenerateSchema[PdfInfoInput](),
			Annotations: &mcp.ToolAnnotations{
				Title:           "PDF Info",
				ReadOnlyHint:    true,
				IdempotentHint:  true,
				DestructiveHint: new(false),
				OpenWorldHint:   new(true),
			},
		},
		func(ctx context.Context, request *mcp.CallToolRequest, input PdfInfoInput) (*mcp.CallToolResult, any, error) {
			args, err := json.Marshal(input)
			if err != nil {
				return nil, nil, fmt.Errorf("marshaling input: %w", err)
			}
			result, err := localServer.ExecuteTool("pdf_info", args)
			if err != nil {
				return nil, nil, fmt.Errorf("error: %w", err)
			}
			return &mcp.CallToolResult{
				Content: []mcp.Content{&mcp.TextContent{Text: result.Content}},
			}, nil, nil
		},
	)
}
