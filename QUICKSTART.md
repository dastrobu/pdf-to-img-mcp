# Quick Start Guide

## Get pdf2img up and running in 5 minutes

### Installation

```bash
# Requires Go 1.27+
go install .

# Done! Use:
pdf2img -i your-document.pdf -o ./output
```

## Getting started

### Convert a complete PDF

```bash
pdf2img -i document.pdf -o ./output
```

**Result:**
- ✓ Creates the `./output` directory
- ✓ Generates `page_0001.png`, `page_0002.png`, etc.

### Get PDF information

```bash
pdf2img info document.pdf
```

**Result:**
```
PDF Information
===============
File: document.pdf
Pages: 42
Size: 2.50 MB
```

### Convert only the first 5 pages

```bash
pdf2img -i document.pdf -o ./output --start 1 --end 5
```

### Use JPG instead of PNG

```bash
pdf2img -i document.pdf -o ./output -f jpg
```

### Increase quality (higher DPI)

```bash
pdf2img -i document.pdf -o ./output -d 300
```

| DPI | Use |
|-----|-----|
| 72-96 | Web, screen |
| 150 | Quality/speed balance (default) |
| 300 | Print, documents |
| 600+ | Highly detailed files |

## Complete options reference

### All available flags

| Flag | Short | Description | Default | Example |
|------|-------|-------------|---------|---------|
| `--input` | `-i` | PDF file (required) | - | `-i document.pdf` |
| `--output` | `-o` | Output directory | `.` (current) | `-o ./output` |
| `--format` | `-f` | Format: `png` or `jpg` | `png` | `-f jpg` |
| `--dpi` | `-d` | DPI for rendering | `150` | `-d 300` |
| `--start` | - | Start page (1-indexed) | `0` (first) | `--start 1` |
| `--end` | - | End page (1-indexed) | `0` (last) | `--end 10` |
| `--prefix` | - | Prefix for files | `page_` | `--prefix img_` |
| `--verbose` | `-v` | Detailed output | `false` | `-v` |
| `--retry` | - | Retry failed pages with reduced DPI | `false` | `--retry` |
| `--pool-size` | - | Max PDFium instances in pool (for large PDFs) | `2` | `--pool-size 4` |
| `--refresh-every` | - | Refresh the WASM instance every N pages (0=disable) | `50` | `--refresh-every 25` |

### Common use cases

#### First-page thumbnail
```bash
pdf2img -i document.pdf -o ./thumbs --start 1 --end 1 -d 96 --prefix thumb_
```

#### Create a web gallery
```bash
pdf2img -i document.pdf -o ./web/images -f jpg -d 150
```

#### Extract a specific range
```bash
pdf2img -i libro.pdf -o ./chapter5 --start 100 --end 150
```

## Option combinations for different scenarios

### 📱 For Web/Screen (low quality, fast)
```bash
pdf2img -i document.pdf -o ./output -f jpg -d 96
```
- Format: JPG (smaller)
- DPI: 96 (sufficient for screen)
- Fast and small files

### 📄 Standard documents (balanced)
```bash
pdf2img -i document.pdf -o ./output -f png -d 150
```
- Format: PNG (lossless)
- DPI: 150 (default, quality/speed balance)
- Good for most cases

### 🖨️ For print/high quality
```bash
pdf2img -i document.pdf -o ./output -f png -d 300
```
- Format: PNG (lossless)
- DPI: 300 (high quality)
- Slower but excellent quality

### 📚 Very large PDFs (>100 pages) - Optimized
```bash
pdf2img -i large-document.pdf -o ./output -d 150 --refresh-every 50 --pool-size 3
```
- `--refresh-every 50`: Refreshes WASM every 50 pages (cleans memory)
- `--pool-size 3`: More PDFium instances for better performance
- Avoids rendering corruption after many pages

### 📚 Very large PDFs (>200 pages) - Maximum control
```bash
pdf2img -i huge-document.pdf -o ./output -d 150 --refresh-every 25 --pool-size 4 --retry
```
- `--refresh-every 25`: Refreshes more frequently (more memory)
- `--pool-size 4`: Maximum instances for better performance
- `--retry`: Retries failed pages with reduced DPI

### ⚡ Fast processing (low quality)
```bash
pdf2img -i document.pdf -o ./output -d 72
```
- DPI: 72 (minimum, very fast)
- Useful for quick previews

### 🎯 Specific pages (range)
```bash
pdf2img -i document.pdf -o ./output --start 50 --end 150
```
- Only converts pages 50 to 150
- Useful for processing very large documents in parts

### 🔍 Specific pages + high quality
```bash
pdf2img -i document.pdf -o ./output --start 1 --end 10 -d 300 -f png
```
- First 10 pages at maximum quality
- Lossless PNG to preserve details

### 📦 Thumbnails
```bash
pdf2img -i document.pdf -o ./thumbs --start 1 --end 1 -d 96 --prefix thumb_ -f jpg
```
- Only the first page
- 96 DPI (good for thumbnails)
- JPG (smaller)
- Prefix: `thumb_`

### 🎬 Image gallery
```bash
pdf2img -i document.pdf -o ./gallery -f jpg -d 150 --prefix gallery_
```
- All pages
- JPG (smaller size)
- 150 DPI (balanced)
- Custom prefix

### 📊 Reports/Analysis
```bash
pdf2img -i report.pdf -o ./analysis -d 200 -f png --prefix report_
```
- Higher than default (200 vs 150 DPI)
- PNG to preserve quality
- Better for detailed analysis

### 🖼️ Detailed files (graphics, blueprints)
```bash
pdf2img -i plano.pdf -o ./output -d 600 -f png
```
- 600 DPI (very high quality)
- PNG (lossless)
- Slower, but excellent for details

### 🧪 Debugging/Diagnosis
```bash
pdf2img -i document.pdf -o ./output -v --retry
```
- `-v`: Detailed output (verbose)
- `--retry`: Retry failed pages
- Useful for diagnosing issues

### 📋 PDF information (no conversion)
```bash
pdf2img info document.pdf
```
- Only shows: pages, size, dimensions
- Does not convert anything

## Advanced combinations for problematic PDFs

### Document with many pages and possible errors
```bash
pdf2img -i document.pdf -o ./output -d 150 --refresh-every 25 --pool-size 4 --retry -v
```
Combines:
- `--refresh-every 25`: Refreshes WASM frequently
- `--pool-size 4`: More resources
- `--retry`: Retries with reduced DPI
- `-v`: See what happens in detail

### Process in blocks (for 300+ page PDFs)
```bash
# Block 1
pdf2img -i document.pdf -o ./output --start 1 --end 75 -d 150 --refresh-every 50

# Block 2
pdf2img -i document.pdf -o ./output --start 76 --end 150 -d 150 --refresh-every 50

# Block 3
pdf2img -i document.pdf -o ./output --start 151 --end 225 -d 150 --refresh-every 50

# Block 4
pdf2img -i document.pdf -o ./output --start 226 --end 300 -d 150 --refresh-every 50
```
Processing in blocks reduces memory issues with very large PDFs.

### Large PDF with variable quality
```bash
# Normal pages with standard DPI
pdf2img -i document.pdf -o ./output --start 1 --end 50 -d 150

# Pages with graphics at higher DPI
pdf2img -i document.pdf -o ./output --start 51 --end 100 -d 300
```

## Performance tips

### If it's slow:
```bash
# Reduce DPI
pdf2img -i document.pdf -o ./output -d 96

# Or specify only the needed range
pdf2img -i document.pdf -o ./output --start 1 --end 10
```

### If pages fail:
```bash
# Use automatic retry
pdf2img -i document.pdf -o ./output --retry

# Or increase refresh
pdf2img -i document.pdf -o ./output --refresh-every 25
```

### If it uses too much memory:
```bash
# Reduce pool size and increase refresh
pdf2img -i document.pdf -o ./output --pool-size 1 --refresh-every 25
```

## Use as an MCP server

### HTTP option (port 8080)

```bash
mcp-server
# Server at: http://localhost:8080

# In another terminal:
curl http://localhost:8080/tools
```

### stdio option (integration with your MCP client)

```bash
mcp-server --stdio
```

## Complete documentation

- [README.md](README.md) - Complete documentation
- [EXAMPLES.md](EXAMPLES.md) - Detailed examples
- [DEVELOPMENT.md](DEVELOPMENT.md) - Developer guide

## Common problems

| Problem | Solution |
|----------|----------|
| `command not found: pdf2img` | Make sure you're in the project directory or install globally with `--global` |
| PDFium not downloading | Run `go mod download` |
| Images look blurry | Increase DPI: `-d 300` |
| The process is slow | Reduce DPI: `-d 96` or specify a range: `--start 1 --end 10` |
| `file not found` | Verify the PDF exists: `ls document.pdf` |

## Next steps

1. ✅ **Getting started**: Convert your first PDF
2. 📖 **Learn more**: Read [EXAMPLES.md](EXAMPLES.md)
3. 🔧 **Customize**: Adjust DPI, format, and options
4. 🌐 **Integrate**: Use as an MCP server or in your Go code
5. 🚀 **Contribute**: PRs are welcome!

## Help

```bash
pdf2img --help
pdf2img info --help
mcp-server --help
```

---

**Questions?** See [README.md](README.md) for complete documentation.
