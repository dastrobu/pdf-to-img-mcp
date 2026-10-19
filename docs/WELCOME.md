# 🎉 Welcome to pdf2img!

## What is pdf2img?

**pdf2img** is a complete Go tool for converting PDF pages to PNG or JPG images. It includes both a **CLI interface** and an **MCP server** for integration with your MCP client (for example Cursor, Zed, or Cline) and other systems.

- ✅ **High-quality rendering** using go-pdfium (based on Google's PDFium)
- 💻 **Complete CLI** with multiple options
- 🔌 **MCP server** for AI integration
- 📚 **Exhaustive documentation**
- 🚀 **Production ready**

## 🚀 Get started in 5 minutes

### Installation
```bash
# Requires Go 1.27+
go install .
pdf2img -i document.pdf -o ./output
```

**Done!** Your images are in `./output/`

## 📚 Documentation

### Quick start
- **[QUICKSTART.md](QUICKSTART.md)** - First steps in 5 minutes

### Learn to use it
- **[README.md](README.md)** - Complete documentation
- **[EXAMPLES.md](EXAMPLES.md)** - Practical examples
- **[INDEX.md](INDEX.md)** - Navigation index

### For developers
- **[PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)** - Project architecture
- **[DEVELOPMENT.md](DEVELOPMENT.md)** - Development guide
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - How to contribute

### Reference
- **[SUMMARY.md](SUMMARY.md)** - Project summary
- **[LICENSE](LICENSE)** - Apache 2.0

## 💡 Quick examples

### Convert all pages
```bash
pdf2img -i document.pdf -o ./output
```

### High resolution (300 DPI)
```bash
pdf2img -i document.pdf -o ./output -d 300
```

### JPG format
```bash
pdf2img -i document.pdf -o ./output -f jpg
```

### First 5 pages only
```bash
pdf2img -i document.pdf -o ./output --start 1 --end 5
```

### PDF information
```bash
pdf2img info document.pdf
```

## 🔌 Use as an MCP server

```bash
# HTTP mode
mcp-server

# stdio mode (your MCP client)
mcp-server --stdio
```

## 🎯 Key features

| Feature | Details |
|---|---|
| **Formats** | PNG, JPG |
| **DPI** | 72 to 600+ configurable |
| **Page range** | Select a specific range |
| **CLI** | 10+ flags and options |
| **MCP Server** | HTTP and stdio |
| **Tests** | Unit tests included |
| **License** | Apache 2.0 |
| **Platforms** | Windows, Linux, macOS |

## 🛠️ Technologies

- **Go 1.21+** - Programming language
- **go-pdfium** - PDF rendering
- **Cobra** - CLI framework
- **PDFium** - Rendering engine

## 📞 Need help?

### Where to start
1. Read [QUICKSTART.md](QUICKSTART.md) if it's your first use
2. Install with the included scripts
3. Try it with your first PDF
4. Explore [EXAMPLES.md](EXAMPLES.md) for more cases

### Common problems
- See [README.md#troubleshooting](README.md#troubleshooting)
- See [EXAMPLES.md#troubleshooting-common-problems](EXAMPLES.md#troubleshooting-common-problems)

### Development
- Read [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)
- Read [DEVELOPMENT.md](DEVELOPMENT.md)
- Run `make test` to verify

### Contribution
- See [CONTRIBUTING.md](CONTRIBUTING.md)
- Follow the style guides
- Create a PR

## ✨ What pdf2img does

### Before (PDF)
```
document.pdf (25 pages)
```

### After (Images)
```
page_0001.png
page_0002.png
...
page_0025.png
```

## 🚀 Use cases

✅ **Convert documents** to images for storage
✅ **Generate thumbnails** for previews
✅ **Extract specific pages** from a PDF
✅ **High resolution for printing** with 300+ DPI
✅ **Integration with your MCP client** via MCP
✅ **Automate processes** with scripts

## 📊 Project structure

```
pdf2img/
├── cmd/pdf2img/            → CLI
├── cmd/mcp-server/         → MCP server
├── pkg/converter/          → Shared logic
├── Documentation (8 .md)
└── Configuration
```

## 🎓 What you'll learn

- How to use go-pdfium to render PDFs
- How to create professional CLIs with Cobra
- How to implement MCP servers
- Best practices in Go
- Clean software architecture

## 🏆 Highlights

- 📚 **Complete documentation** (8 guides)
- 🧪 **Unit tests** included
- 🔧 **Installation scripts** (Windows + Unix)
- ✨ **Clean code** and professional
- 🚀 **Production ready**
- 📜 **Free license** (Apache 2.0)

## 📈 Roadmap

- [ ] Support for password-protected PDFs
- [ ] Parallel processing
- [ ] Image OCR
- [ ] Rendering cache
- [ ] More formats (WebP, TIFF)

## 🎯 Next step

**Ready to start?**

→ [QUICKSTART.md](QUICKSTART.md) (5 minutes)

**Want to understand everything?**

→ [INDEX.md](INDEX.md) (Complete navigation)

**Want the summary?**

→ [SUMMARY.md](SUMMARY.md) (Project summary)

---

## 📝 Project information

| Aspect | Details |
|---|---|
| **Version** | 1.0.0 |
| **License** | Apache 2.0 |
| **Status** | Complete and ready |
| **Platforms** | Windows, Linux, macOS |
| **Go version** | 1.21+ |
| **Dependencies** | go-pdfium, Cobra |

---

## 🎉 Get started now!

1. **Read** [QUICKSTART.md](QUICKSTART.md) (5 min)
2. **Install** with `go install .`
3. **Try it** with `pdf2img -i document.pdf -o ./output`
4. **Explore** [EXAMPLES.md](EXAMPLES.md) for more

---

**Created with ❤️ using Go and PDFium**

Enjoy converting your PDFs to images! 📸

Questions? See [INDEX.md](INDEX.md) for complete navigation.
