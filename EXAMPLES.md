# Usage Examples

## CLI - Practical examples

### 1. Basic conversion - All pages to PNG

```bash
pdf2img -i document.pdf -o ./images
```

Result:
```
✓ Conversion Complete
Total pages: 42
Successful: 42
Failed: 0
```

Generated files:
- `images/page_0001.png`
- `images/page_0002.png`
- ...
- `images/page_0042.png`

### 2. Convert to JPG with high resolution

```bash
pdf2img -i document.pdf -o ./output -f jpg -d 300
```

For printing, use DPI 300. For screen, 72-150 is sufficient.

### 3. Extract a specific page range

```bash
pdf2img -i libro.pdf -o ./thumbs --start 1 --end 10
```

Only converts the first 10 pages.

### 4. Use a custom prefix

```bash
pdf2img -i report.pdf -o ./output --prefix img_
```

Generates: `img_0001.png`, `img_0002.png`, etc.

### 5. Conversion with detailed output

```bash
pdf2img -i document.pdf -o ./output -v
```

Shows each generated file and detailed errors.

### 6. Get PDF information

```bash
pdf2img info document.pdf
```

Output:
```
PDF Information
===============
File: document.pdf
Pages: 25
Size: 2.50 MB
Width: 612.00 pt
Height: 792.00 pt
```

## MCP Server - Integration examples

### 1. Configuration with your MCP client (Recommended)

#### Requirements
- A compatible MCP client (for example Cursor, Zed, or Cline)
- `mcp-server` compiled in the project

#### Step 1: Locate the configuration file

Each MCP client stores its configuration in a JSON file (usually `mcp.json` or similar). The exact location depends on the client; consult your MCP client's documentation.

As a reference, on a typical system the file is usually at:

**macOS**:
```
~/Library/Application Support/<tu-cliente>/mcp.json
```

**Windows**:
```
%APPDATA%\<tu-cliente>\mcp.json
```

**Linux**:
```
~/.config/<tu-cliente>/mcp.json
```

#### Step 2: Add the pdf2img configuration

Edit your MCP client's configuration file and add the following section under `mcpServers`:

```json
{
  "mcpServers": {
    "pdf2img": {
      "command": "/ruta/al/proyecto/mcp-server",
      "args": ["--stdio"],
      "env": {}
    }
  }
}
```

**Important**: Use the full path where you compiled `mcp-server`. For example:
```json
{
  "mcpServers": {
    "pdf2img": {
      "command": "/home/user/projects/mcp-go-pdf-to-img/mcp-server",
      "args": ["--stdio"],
      "env": {}
    }
  }
}
```

#### Step 3: Restart your MCP client

Close and reopen your MCP client. You should see the "pdf2img" MCP server connected.

#### Step 4: Use with your MCP client

You can now use pdf2img directly from your MCP client:

**Example 1**: "Convert document.pdf to PNG images at 300 DPI"

Your AI assistant will automatically use the `pdf_to_images` tool with the correct parameters.

**Example 2**: "How many pages does ejemplo.pdf have?"

Your AI assistant will use `pdf_info` to get the information.

### 2. Available Tools

#### Tool: `pdf_to_images`

Converts pages of a PDF to PNG or JPG images.

**Parameters**:
```json
{
   "pdf_path": "document.pdf",             // PDF path (required)
   "output_dir": "./output",                // Output directory (required)
   "format": "png",                         // png or jpg (optional, default: png)
   "dpi": 150,                              // Resolution (optional, default: 150)
   "start_page": 1,                         // First page (optional, default: all)
   "end_page": 10,                          // Last page (optional, default: all)
   "prefix": "page_"                        // File prefix (optional, default: page_)
}
```

**Response**:
```json
{
   "total_pages": 42,
   "successful": 42,
   "failed": 0,
   "files": [
     "./output/page_0001.png",
     "./output/page_0002.png",
     ...
   ]
}
```

#### Tool: `pdf_info`

Gets information about a PDF.

**Parameters**:
```json
{
   "pdf_path": "document.pdf"    // PDF path (required)
}
```

**Response**:
```json
{
   "file": "document.pdf",
   "pages": 25,
   "file_size": "2.50 MB",
   "width": 612.00,
   "height": 792.00
}
```

#### Tool: `pdf_compress`

Compresses a PDF to reduce its size by optimizing images and removing unnecessary elements.

**Parameters**:
```json
{
   "pdf_path": "document.pdf",          // Path of the PDF to compress (required)
   "output_path": "compressed-document.pdf"    // Path of the compressed PDF (required)
}
```

**Response**:
```
PDF compressed successfully. Output saved to: compressed-document.pdf
```

### 3. HTTP mode (For testing)

If you prefer to use HTTP instead of stdio mode with your MCP client:

```bash
mcp-server --port 8080
```

#### Get available tools

```bash
curl http://localhost:8080/tools
```

Response:
```json
[
  {
    "name": "pdf_to_images",
    "description": "Convert PDF pages to PNG or JPG images",
    "inputSchema": { ... }
  },
  {
    "name": "pdf_info",
    "description": "Get information about a PDF file",
    "inputSchema": { ... }
  }
]
```

#### Convert PDF to images

```bash
curl -X POST http://localhost:8080/execute \
  -H "Content-Type: application/json" \
  -d '{
    "tool": "pdf_to_images",
    "input": {
      "pdf_path": "document.pdf",
      "output_dir": "./output",
      "format": "png",
      "dpi": 150
    }
  }'
```

Response:
```json
{
   "type": "text",
  "content": "{\n  \"total_pages\": 42,\n  \"successful\": 42,\n  \"failed\": 0,\n  \"files\": [\n    \"./output/page_0001.png\",\n    \"./output/page_0002.png\",\n    ...\n  ]\n}\n"
}
```

#### Get PDF information

```bash
curl -X POST http://localhost:8080/execute \
  -H "Content-Type: application/json" \
  -d '{
    "tool": "pdf_info",
    "input": {
      "pdf_path": "document.pdf"
    }
  }'
```

### 4. stdio mode (For integration in scripts)

```bash
mcp-server --stdio
```

The server reads JSON commands from stdin and writes results to stdout:

```bash
echo '{"tool": "pdf_info", "input": {"pdf_path": "document.pdf"}}' | mcp-server --stdio
```

## From Go code

### Basic

```go
package main

import (
	"log"
	"github.com/scopweb/mcp-go-pdf-to-img/pkg/converter"
)

func main() {
	// Create converter
	conv, err := converter.New()
	if err != nil {
		log.Fatal(err)
	}
	defer conv.Close()

	// Convert
	result, err := conv.Convert(&converter.ConvertOptions{
		InputPath: "document.pdf",
		OutputDir: "./output",
		Format:    "png",
		DPI:       150,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successful: %d, Failed: %d", result.Successful, result.Failed)
}
```

### With MCP Server

```go
package main

import (
	"encoding/json"
	"log"
	"github.com/scopweb/mcp-go-pdf-to-img/mcp"
)

func main() {
	server, err := mcp.NewMCPServer()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	// Convert PDF using MCP
	input := json.RawMessage(`{
		"pdf_path": "document.pdf",
		"output_dir": "./output",
		"format": "png",
		"dpi": 150
	}`)

	result, err := server.ExecuteTool("pdf_to_images", input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result.Content)
}
```

## Real-world use cases

### 1. Generate PDF thumbnails in batch

```bash
#!/bin/bash
for pdf in *.pdf; do
    output_dir="thumbnails/${pdf%.pdf}"
    mkdir -p "$output_dir"
    pdf2img -i "$pdf" -o "$output_dir" -f jpg -d 96 --start 1 --end 1
    echo "✓ Thumbnail of $pdf generated"
done
```

### 2. Convert PDF to a web image gallery

```bash
pdf2img -i articulo.pdf -o ./web/images -f jpg -d 150
# Then create HTML that lists the images
for img in web/images/*.jpg; do
    echo "<img src='$(basename $img)' />"
done > web/gallery.html
```

### 3. Integration with a web service

```go
// In your web application
func handlePDFUpload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Upload failed", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save PDF temporarily
	tmpFile, _ := ioutil.TempFile("", "*.pdf")
	io.Copy(tmpFile, file)
	tmpFile.Close()

	// Convert
	conv, _ := converter.New()
	defer conv.Close()
	result, _ := conv.Convert(&converter.ConvertOptions{
		InputPath: tmpFile.Name(),
		OutputDir: "./static/uploads",
		Format:    "jpg",
		DPI:       150,
	})

	// Return image URLs
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.OutputFiles)
}
```

### 4. Background processing with workers

```go
type PDFJob struct {
	Path   string
	Output string
	DPI    float64
}

func worker(jobs <-chan PDFJob, results chan<- PDFJob) {
	conv, _ := converter.New()
	defer conv.Close()

	for job := range jobs {
		conv.Convert(&converter.ConvertOptions{
			InputPath: job.Path,
			OutputDir: job.Output,
			DPI:       job.DPI,
		})
		results <- job
	}
}

// Create worker pool
jobs := make(chan PDFJob, 100)
results := make(chan PDFJob)

for i := 0; i < 4; i++ {
	go worker(jobs, results)
}

// Enqueue jobs
for _, pdf := range pdfFiles {
	jobs <- PDFJob{Path: pdf, Output: "./output", DPI: 150}
}
```

## Troubleshooting common problems

### Images look blurry
Increase DPI:
```bash
pdf2img -i document.pdf -o ./output -d 300
```

### The process is very slow
Reduce DPI:
```bash
pdf2img -i document.pdf -o ./output -d 96
```

### Out of Memory
Process by ranges:
```bash
# Pages 1-100
pdf2img -i document.pdf -o ./out1 --start 1 --end 100

# Pages 101-200
pdf2img -i document.pdf -o ./out2 --start 101 --end 200
```

### Error: "PDF file not found"
Verify the path:
```bash
ls -la document.pdf
# Use an absolute path
pdf2img -i /path/complete/document.pdf -o ./output
```

### MCP server not responding
Verify the port is available:
```bash
netstat -tuln | grep 8080
# Or use another port
mcp-server -port 9090
```
