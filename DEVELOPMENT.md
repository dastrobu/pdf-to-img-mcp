# Development Guide

## Development environment setup

### Requirements
- Go 1.21+
- Git
- Make (optional, but recommended)

### Initial setup

```bash
# Clone the repository
git clone <repository-url>
cd pdf2img

# Download dependencies
go mod download
go mod verify

# Build the binary
make build
# or
go build -o pdf2img .
```

## Code structure

### `pkg/converter/converter.go`
Main module that encapsulates the PDF conversion logic:
- `Converter`: Main structure that manages the PDFium instance
- `ConvertOptions`: Options to configure the conversion
- `ConvertResult`: Result of the conversion
- Helper functions for validation, page range, and saving images

### `main.go` / `split.go`
CLI application using Cobra (the `main` package at the root, installable with `go install .`):
- Main command `convert` (by default)
- `info` subcommand to get PDF metadata
- Flag and command-line option handling
- Output formatting for the user

### `mcp/server.go`
Model Context Protocol server implementation:
- `MCPServer`: Structure that exposes MCP tools
- `GetTools()`: Lists available tools
- `ExecuteTool()`: Executes a tool with JSON input
- Handlers for `pdf_to_images` and `pdf_info`

### `cmd/mcp-server/main.go`
HTTP/stdio server for MCP:
- HTTP mode: REST API on localhost
- stdio mode: For integration with any MCP client

## Compilation and testing

### Compile
```bash
# Optimized compilation (release)
make build

# Compilation with debug symbols
make build-dev

# CLI-specific compilation
go build -o pdf2img .

# MCP server-specific compilation
go build -o mcp-server ./cmd/mcp-server
```

### Tests
```bash
# Run all tests
make test
# or
go test -v ./...

# Tests for a specific package
go test -v ./pkg/converter
```

### Linting and formatting
```bash
# Format code
make fmt

# Lint (requires golangci-lint installed)
make lint
```

## Working with go-pdfium

### Initialization
```go
conv, err := converter.New()
defer conv.Close()
```

### Get PDF information
```go
info, err := conv.GetPDFInfo("document.pdf")
// info["pages"], info["width"], info["height"], etc.
```

### Convert PDF
```go
result, err := conv.Convert(&converter.ConvertOptions{
    InputPath: "document.pdf",
    OutputDir: "./output",
    Format:    "png",
    DPI:       150,
    Prefix:    "page_",
})
```

## Future extensions

### 1. Support for password-protected PDFs
```go
// In converter.go
doc, err := c.pdfium.OpenFile(pdfPath, &requests.OpenFile{
    Password: "your-password",
})
```

### 2. Parallel processing
```go
// Render multiple pages concurrently
var wg sync.WaitGroup
for pageNum := startPage; pageNum <= endPage; pageNum++ {
    wg.Add(1)
    go func(p int) {
        defer wg.Done()
        // Render page
    }(pageNum)
}
wg.Wait()
```

### 3. Complete REST API
Extend `cmd/mcp-server/main.go` with more endpoints:
```go
http.HandleFunc("/api/v1/convert", handleConvert)
http.HandleFunc("/api/v1/info", handleInfo)
http.HandleFunc("/api/v1/status", handleStatus)
```

### 4. OCR of images
Integrate with `github.com/otiai10/gosseract`:
```go
client, _ := gosseract.NewClient()
defer client.Close()
text, _ := client.Text()
```

### 5. Rendering cache
Store generated images to avoid re-rendering:
```go
type Cache struct {
    mu    sync.RWMutex
    items map[string][]byte
}
```

## Debugging

### With detailed logs
```bash
pdf2img -i document.pdf -o ./output -v
```

### In code
```go
import "log"

log.Printf("Debug: %+v", variable)
```

### Go profiling
```bash
# CPU profiling
go run -cpuprofile=cpu.prof .

# Memory profiling
go run -memprofile=mem.prof .

# Analysis
go tool pprof cpu.prof
```

## Dependencies

| Library | Version | Purpose | License |
|----------|---------|---------|----------|
| go-pdfium | v1.12.0 | Render PDFs | Apache 2.0 |
| Cobra | v1.7.0 | CLI framework | Apache 2.0 |

To update dependencies:
```bash
go get -u github.com/klippa-app/go-pdfium
go get -u github.com/spf13/cobra
go mod tidy
```

## Contribution

1. Fork the project
2. Create a branch (`git checkout -b feature/AmazingFeature`)
3. Make clear commits (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Apache 2.0 - See LICENSE
