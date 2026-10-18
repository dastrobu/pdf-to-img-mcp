#!/bin/bash
# Installation script for pdf2img

set -e

echo "📦 Installing pdf2img..."

# Check for Go
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    echo "   Visit: https://golang.org/dl"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
echo "✓ Found Go version: $GO_VERSION"

# Download dependencies
echo "📥 Downloading dependencies..."
go mod download
go mod verify

# Build CLI
echo "🔨 Building CLI..."
go build -o pdf2img .
echo "✓ CLI built: ./pdf2img"

# Build MCP server (optional)
if [ "$1" = "--with-mcp" ]; then
    echo "🔨 Building MCP server..."
    go build -o mcp-server ./cmd/mcp-server
    echo "✓ MCP server built: ./mcp-server"
fi

# Test installation
echo "🧪 Testing installation..."
./pdf2img --help > /dev/null && echo "✓ CLI works correctly"

# Optional: Install globally
if [ "$1" = "--global" ] || [ "$2" = "--global" ]; then
    echo "🌐 Installing globally..."
    go install .
    echo "✓ Installed globally as 'mcp-go-pdf-to-img' (build with an explicit name via: go build -o pdf2img .)"

    if [ "$1" = "--with-mcp" ] || [ "$2" = "--with-mcp" ]; then
        go install ./cmd/mcp-server
        echo "✓ Installed globally as 'mcp-server'"
    fi
fi

echo ""
echo "✅ Installation complete!"
echo ""
echo "Usage:"
echo "  ./pdf2img -i input.pdf -o ./output"
echo "  ./pdf2img info input.pdf"
echo ""
echo "For more information, see README.md"
