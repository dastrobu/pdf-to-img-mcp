# AGENTS.md

Generic, tool-agnostic instructions for AI coding agents working in this repository.

## Overview

`mcp-go-pdf-to-img` converts PDF pages to PNG/JPG images using [go-pdfium](https://github.com/klippa-app/go-pdfium) (PDFium compiled to WebAssembly — pure Go, no CGO). It ships three binaries plus a reusable library:

- **CLI** — a command-line tool (`pdf2img`) for converting, inspecting, and splitting PDFs.
- **MCP server** — a stdio [Model Context Protocol](https://modelcontextprotocol.io/) server exposing conversion/info tools to any MCP client.
- **HTTP server** — a REST API microservice for PDF→image conversion.
- **Library** — `pkg/converter` (rendering) and `pkg/splitter` (page extraction), usable directly from Go code.

## Prerequisites

- **Go 1.27+** (see `go.mod`, `go 1.27`).
- **No external dependencies** — everything is embedded in the binary. Rendering runs in a WebAssembly sandbox via wazero, so there is no CGO, no C toolchain, and no system-level PDFium to install. A single `go build` produces a portable, cross-platform binary.

## Build

```bash
go build ./...                 # build everything
go build -o pdf2img .          # CLI (root `main` package)
go build -o mcp-server ./cmd/mcp-server   # stdio MCP server
go build -o server ./cmd/server           # HTTP REST server
```

`go install .` installs the CLI (the root `main` package). Note that installing without an explicit name produces a binary named after the module; use `go build -o pdf2img .` to get a `pdf2img` binary.

## Test

```bash
go test ./...      # run the full test suite
gofmt -l .         # should list nothing (all files formatted)
go vet ./...       # should pass cleanly
```

Tests live in:
- `pkg/converter/converter_test.go` — unit tests for conversion logic.
- `cmd/server/handlers_test.go` — HTTP handler tests.
- `test/security/` — security/CVE regression tests.

## Run

**CLI**

```bash
# Convert a PDF to images
pdf2img -i doc.pdf -o out [-f png|jpg -d 150 --start N --end M --prefix p_ --retry]

# Inspect a PDF
pdf2img info doc.pdf

# Extract a page range into a new PDF
pdf2img split -i doc.pdf -o out.pdf --start N --end M
```

Key flags: `-i`/`--input` (required), `-o`/`--output`, `-f`/`--format` (png|jpg, default png), `-d`/`--dpi` (default 150), `--start`/`--end` (1-indexed, 0 = first/last), `--prefix` (default `page_`), `-v`/`--verbose`, `--retry`, plus `--pool-size` and `--refresh-every` for large PDFs.

**MCP server (stdio)**

```bash
mcp-server --stdio
```

**HTTP server**

```bash
server          # defaults: 0.0.0.0:8080
HTTP_PORT=8080 HTTP_HOST=0.0.0.0 CONVERT_DEFAULT_DPI=200 server
```

Configuration is via environment variables, e.g. `HTTP_PORT`, `HTTP_HOST`, `HTTP_READ_TIMEOUT`, `HTTP_WRITE_TIMEOUT`, `HTTP_MAX_UPLOAD_SIZE`, `LOG_LEVEL`, `LOG_FORMAT`, `CONVERT_DEFAULT_DPI`, `CONVERT_DEFAULT_FORMAT`, `CONVERT_DEFAULT_POOL_SIZE`, `CONVERT_REFRESH_EVERY`, `CONVERT_RETRY_ON_FAILURE`.

## Project Structure

```
.
├── main.go                  # CLI root command (package main)
├── split.go                 # `split` subcommand
├── cmd/
│   ├── mcp-server/main.go   # stdio MCP server (tools: pdf_to_images, pdf_info)
│   └── server/              # HTTP REST server (main.go, handlers.go)
├── pkg/
│   ├── converter/converter.go    # PDFium (WASM) rendering, chunked to limit WASM state
│   └── splitter/splitter.go      # page extraction via pdfcpu
├── mcp/
│   ├── server.go             # local MCP implementation (pdf_to_images, pdf_info, pdf_compress, pdf_split)
│   └── example_server.go     # usage example
├── internal/
│   ├── config/config.go      # HTTP server + conversion config from env vars
│   └── logging/logger.go     # slog-based logger
└── test/security/            # security regression tests
```

- **CLI** (`main.go`, `split.go`) — `package main`; the root command converts PDFs, with `info` and `split` subcommands (cobra).
- **`cmd/mcp-server`** — stdio MCP server built on the `modelcontextprotocol/go-sdk` (official MCP Go SDK); exposes `pdf_to_images` and `pdf_info`.
- **`cmd/server`** — HTTP REST server with `GET /health`, `POST /api/v1/pdf/convert`, `POST /api/v1/pdf/info`.
- **`pkg/converter`** — PDFium rendering via WebAssembly; pages are processed in chunks (default every 50, via `--refresh-every`) so the WASM instance is refreshed to avoid state accumulation on large PDFs.
- **`pkg/splitter`** — page-range extraction using `pdfcpu`.
- **`mcp/`** — a local, dependency-free MCP implementation that also exposes `pdf_compress` and `pdf_split`.

## Conventions

- Standard Go style; run `gofmt` and `go vet` before committing.
- Lowercase package names, no underscores (e.g. `converter`, `splitter`, `config`).
- Exported identifiers `PascalCase`; unexported `camelCase`.
- Constants in `SCREAMING_SNAKE_CASE`.
- Wrap errors with `fmt.Errorf("...: %w", err)` to preserve the chain.
- Always release resources with `defer` — e.g. `defer conv.Close()` for the converter, and `defer f.Close()` for files.
- Avoid `panic`; it should only appear, if at all, during initialization. Return errors instead.
- Table-driven tests named `TestXxx` (e.g. `TestValidateOptions`, `TestDeterminePageRange`, `TestHealth`).

## MCP Server Note

The MCP server is client-agnostic — it speaks the Model Context Protocol over stdio and works with any MCP client (Cursor, Zed, Cline, and others). See `MCP.md` for client integration details.
