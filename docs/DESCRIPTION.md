cat > DESCRIPTION.md << 'EOF'
# mcp-go-pdf-to-img

A high-performance Go utility for converting PDF pages to PNG/JPG images 
with your MCP client (for example Cursor, Zed, or Cline) integration.

## Overview

This project provides a pure Go implementation for batch PDF to image conversion 
with seamless integration into your MCP client via the Model Context Protocol (MCP).

## Key Features

- **Pure Go + WebAssembly**: No CGO dependencies, fully portable
- **Single Binary**: ~16-18 MB executable for Windows, Linux, macOS
- **CLI Tool**: Command-line interface for batch operations
- **MCP Server**: Direct integration with your MCP client
- **Flexible Output**: PNG/JPG with customizable DPI and page ranges
- **Production Ready**: Based on go-pdfium v1.17.2 with official MCP SDK

## Technology Stack

- Go 1.21+
- go-pdfium v1.17.2 (WebAssembly-based PDF rendering)
- mark3labs/mcp-go v0.43.0 (Official MCP SDK)
- Cobra CLI framework

## Quick Links

- [README.md](README.md) - Complete documentation
- [MCP.md](MCP.md) - MCP client setup
- [EXAMPLES.md](EXAMPLES.md) - Usage examples

## License

Apache 2.0
EOF