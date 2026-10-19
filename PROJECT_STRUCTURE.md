# Project Structure

## Directory tree

```
pdf2img/
├── main.go                         # Application CLI (root main package — `go install .`)
├── split.go                        # `split` subcommand
├── cmd/                            # Additional executable applications
│      ├── mcp-server/                 # MCP server
│      │      └── main.go                # HTTP/stdio server
│      └── server/                     # HTTP REST server
│          ├── main.go
│          └── handlers.go
├── internal/                        # Internal packages
│      ├── config/                   # Server configuration
│      │      └── config.go
│      └── logging/                  # Logging utilities
│          └── logger.go
├── mcp/                            # MCP protocol implementation
│      ├── server.go                  # MCP server
│      └── example_server.go          # Usage example
├── pkg/                            # Reusable packages
│      ├── converter/                 # Conversion logic
│      │      ├── converter.go           # Main logic
│      │      └── converter_test.go      # Unit tests
│      └── splitter/                  # Page extraction
│          └── splitter.go
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums (generated)
├── Makefile                         # Build automation
├── .gitignore                       # Files to ignore in Git
├── LICENSE                         # Apache 2.0 license
├── README.md                        # Main documentation
├── QUICKSTART.md                   # Quick start guide
├── EXAMPLES.md                     # Usage examples
├── DEVELOPMENT.md                  # Developer guide
└── PROJECT_STRUCTURE.md            # This file
```

## Component description

### `main.go` / `split.go`
**Main CLI application** (the `main` package at the root, installable with `go install .`)

- Command-line interface to convert PDFs
- Uses Cobra for flag handling
- Supports multiple options: DPI, format, page range, etc.
- `info` command to get PDF metadata

### `cmd/mcp-server/`
**MCP server (Model Context Protocol)**

- Provides two interfaces:
   - **HTTP**: REST API on port 8080
   - **stdio**: For integration with any MCP client
- Exposes MCP tools:
   - `pdf_to_images`: Convert PDF to images
   - `pdf_info`: Get PDF information

### `mcp/`
**MCP server logic**

- `server.go`: Server implementation
   - `NewMCPServer()`: Create instance
   - `GetTools()`: List available tools
   - `ExecuteTool()`: Execute a tool
   - Handlers for each tool

- `example_server.go`: Programmatic usage example

### `pkg/converter/`
**Shared conversion logic**

- `converter.go`: Main module
   - `New()`: Initialize PDFium
   - `Convert()`: Render PDF to images
   - `GetPDFInfo()`: Get metadata
   - Helper functions for validation and saving

- `converter_test.go`: Unit tests
   - Options validation
   - Page range determination
   - File size formatting

## Data flow

### PDF conversion (CLI)

```
CLI (main.go)
   ├─→ Parse arguments (Cobra)
   ├─→ Validate options
   ├─→ Create Converter
   ├─→ converter.Convert()
   │    ├─→ Validate options
   │    ├─→ PDFium.OpenFile()
   │    ├─→ For each page:
   │    │    ├─→ PDFium.RenderPageInDPI()
   │    │    └─→ saveImage() (PNG/JPG)
   │    └─→ Return ConvertResult
   └─→ Display results
```

### Conversion via MCP

```
Client (HTTP/stdio)
   ├─→ Send JSON {"tool": "pdf_to_images", "input": {...}}
   ├─→ MCPServer.ExecuteTool()
   │    ├─→ converter.Convert()
   │    └─→ Format result as JSON
   └─→ Return JSON result
```

## Dependencies

```
github.com/klippa-app/go-pdfium (v1.12.0)
   └─→ PDFium (C++)
       └─→ PDF rendering

github.com/spf13/cobra (v1.7.0)
   └─→ CLI framework with flags
```

## Configuration

### Environment variables
Currently no environment variables are required.

### Configuration files
- Integration with MCP clients is configured in your MCP client's configuration file (no files are included in this repository).

## Conventions

### Naming
- Packages: `lowercase` without underscore
- Functions/methods: `PascalCase` (public), `camelCase` (private)
- Constants: `SCREAMING_SNAKE_CASE`
- Variables: `camelCase`

### Test structure
- File name: `{name}_test.go`
- Function name: `Test{FunctionToTest}`
- Case table: `tests := []struct {...}`

### Error handling
- Propagate errors with context: `fmt.Errorf("action failed: %w", err)`
- Don't use `panic()` except during initialization
- Always defer for cleanup: `defer converter.Close()`

## Compilation

### Compile the CLI
```bash
go build -o pdf2img .
```

### Compile the MCP server
```bash
go build -o mcp-server ./cmd/mcp-server
```

### Optimized compilation
```bash
go build -ldflags "-s -w" -o pdf2img .
```

## Testing

### Run all tests
```bash
go test -v ./...
```

### Specific test
```bash
go test -v ./pkg/converter -run TestValidateOptions
```

### Coverage
```bash
go test -cover ./...
```

## Performance

### Optimizations implemented
1. **Lazy initialization**: PDFium is initialized on demand
2. **Resource cleanup**: `defer` to close files and free memory
3. **Efficient error handling**: Without recreating unnecessary objects

### Possible improvements
1. Parallel page processing
2. Cache of rendered images
3. Worker pool for batch processing
4. Output image compression

## Security

### Validations
1. Input file verification
2. Output format validation
3. Page range validation
4. Safe directory creation

### Considerations
1. Don't accept absolute paths without validation
2. Sanitize output names
3. Limit maximum DPI
4. Protect against malicious PDFs

## Internationalization (i18n)
Currently no i18n support. All messages in English.

For future expansion:
- Use `golang.org/x/text/language`
- Separate strings into message files
- Translate automatically

## Versioning
Semantic (MAJOR.MINOR.PATCH):
- MAJOR: Incompatible changes
- MINOR: New features
- PATCH: Bug fixes

Currently: v1.0.0

## Roadmap

### v1.1.0
- [ ] Support for password-protected PDFs
- [ ] Parallel page processing
- [ ] Better error reporting

### v1.2.0
- [ ] OCR of images
- [ ] Rendering cache
- [ ] Additional HTTP API

### v2.0.0
- [ ] Architecture redesign
- [ ] Support for more formats (WebP, TIFF)
- [ ] Database for metadata
