# 📑 pdf2img Project Index

## 🚀 Quick start

### Want to get started right away?
→ [QUICKSTART.md](QUICKSTART.md) 

### Want to understand the project?
→ [SUMMARY.md](SUMMARY.md)

### Want the complete documentation?
→ [README.md](README.md)

---

## 📚 Documentation by topic

### For users

| Document | Content |
|-----------|----------|
| [QUICKSTART.md](QUICKSTART.md) | Installation and first steps |
| [README.md](README.md) | Complete documentation and reference |
| [EXAMPLES.md](EXAMPLES.md) | Practical usage examples |

**Recommended reading order:** QUICKSTART → README → EXAMPLES

### For developers

| Document | Content |
|-----------|----------|
| [DEVELOPMENT.md](DEVELOPMENT.md) | Setup, build, debugging |
| [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) | Internal architecture |
| [CONTRIBUTING.md](CONTRIBUTING.md) | How to contribute |

**Recommended reading order:** PROJECT_STRUCTURE → DEVELOPMENT → CONTRIBUTING

### Reference

| Document | Content |
|-----------|----------|
| [LICENSE](LICENSE) | Apache 2.0 License |
| [SUMMARY.md](SUMMARY.md) | Project summary |

---

## 🎯 By use case

### "I need to convert a PDF to images"
1. [QUICKSTART.md](QUICKSTART.md) - Installation
2. Run: `pdf2img -i document.pdf -o ./output`

### "I want to understand how it works"
1. [SUMMARY.md](SUMMARY.md) - General summary
2. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - Architecture
3. Read the code in `pkg/converter/converter.go`

### "I want to use the MCP server with my MCP client"
1. [README.md](README.md#mcp-server) - MCP section
2. [EXAMPLES.md](EXAMPLES.md#mcp-server---integration-examples) - MCP examples

### "I want to contribute to the project"
1. [DEVELOPMENT.md](DEVELOPMENT.md) - Local setup
2. [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guide
3. [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - Understand the code

### "I need specific information"
1. [README.md](README.md#troubleshooting) - Troubleshooting
2. [EXAMPLES.md](EXAMPLES.md#troubleshooting-common-problems) - Common problems

---

## 📂 Directory structure

```
pdf2img/
├── cmd/                           # Applications
│   ├── pdf2img/main.go           # CLI
│   └── mcp-server/main.go        # MCP server
├── mcp/                           # MCP logic
├── pkg/converter/                 # Conversion logic
├── [Documentation]                # 8 .md files
├── go.mod, Makefile, etc.
└── LICENSE
```

See [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) for complete details.

---

## 🔍 Quick search

### Installation
- [QUICKSTART.md](QUICKSTART.md)
- [README.md#installation](README.md#installation)

### CLI usage
- [README.md#usage](README.md#usage)
- [EXAMPLES.md#cli---practical-examples](EXAMPLES.md#cli---practical-examples)

### MCP server
- [README.md#mcp-server](README.md#mcp-server)
- [EXAMPLES.md#mcp-server](EXAMPLES.md#mcp-server---integration-examples)

### Troubleshooting
- [README.md#troubleshooting](README.md#troubleshooting)
- [EXAMPLES.md#troubleshooting-common-problems](EXAMPLES.md#troubleshooting-common-problems)

### Development
- [DEVELOPMENT.md](DEVELOPMENT.md)
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)

### Contribution
- [CONTRIBUTING.md](CONTRIBUTING.md)

---

## 📊 Content by file

### README.md (5.6 KB)
- Key features
- Installation (3 methods)
- CLI usage
- MCP server
- Dependencies
- Roadmap
- Troubleshooting

### QUICKSTART.md (3.2 KB)
- Quick installation (3 options)
- First steps
- Common cases
- Quick problems

### EXAMPLES.md (6.1 KB)
- CLI: 6 practical examples
- MCP HTTP: Complete examples
- From Go code
- Real use cases
- Advanced troubleshooting

### DEVELOPMENT.md (4.4 KB)
- Environment setup
- Code structure
- Build and testing
- Debugging
- Future extensions

### PROJECT_STRUCTURE.md (6.6 KB)
- Directory tree
- Component description
- Data flow
- Dependencies
- Code conventions

### CONTRIBUTING.md (5.7 KB)
- Code of conduct
- How to report bugs
- Submitting PRs
- Style guide
- Review process

### SUMMARY.md (7.4 KB)
- Project summary
- Completed features
- Statistics
- Available documentation
- Next steps

---

## 🎯 Recommended navigation flow

```
First use?
     ├─→ QUICKSTART.md
     ├─→ Install
     ├─→ Try it with a PDF
     └─→ README.md for more details

Use as a developer?
     ├─→ SUMMARY.md
     ├─→ PROJECT_STRUCTURE.md
     ├─→ DEVELOPMENT.md
     └─→ Read the code

Contribute?
     ├─→ PROJECT_STRUCTURE.md
     ├─→ DEVELOPMENT.md
     ├─→ CONTRIBUTING.md
     └─→ Create a PR

Use the MCP server?
     ├─→ README.md (MCP section)
     ├─→ EXAMPLES.md (MCP section)
     └─→ Integrate with your MCP client

Need help?
     ├─→ README.md (Troubleshooting)
     ├─→ EXAMPLES.md (Common problems)
     └─→ Open an issue
```

---

## 📞 Quick navigation

### 🔧 Installation
- [Installation](QUICKSTART.md#installation)

### 💻 Commands
- [Basic CLI](README.md#cli---basic-conversion)
- [Advanced CLI](EXAMPLES.md)
- [MCP server](README.md#mcp-server)

### 🚀 Use cases
- [Thumbnails](EXAMPLES.md#1-generate-pdf-thumbnails-in-batch)
- [Web gallery](EXAMPLES.md#2-convert-pdf-to-a-web-image-gallery)
- [Web integration](EXAMPLES.md#3-integration-with-a-web-service)

### 📚 Concepts
- [Architecture](PROJECT_STRUCTURE.md#data-flow)
- [Conventions](PROJECT_STRUCTURE.md#conventions)
- [Performance](PROJECT_STRUCTURE.md#performance)

---

## 🆘 SOS - Need help

| Problem | Solution |
|----------|----------|
| I don't know where to start | [QUICKSTART.md](QUICKSTART.md) |
| I have an error | [Troubleshooting](README.md#troubleshooting) |
| I want to understand the code | [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) |
| I have a question | [EXAMPLES.md](EXAMPLES.md) or [README.md](README.md) |
| I want to contribute | [CONTRIBUTING.md](CONTRIBUTING.md) |
| The server doesn't work | [DEVELOPMENT.md](DEVELOPMENT.md#debugging) |

---

## 📈 Statistics

- **21** files total
- **6** Go files
- **7** documentation files
- **~600+** lines of code
- **25+** KB of documentation
- **License**: Apache 2.0

---

## ✨ Highlights

- ✅ Complete CLI
- ✅ MCP server
- ✅ Exhaustive documentation
- ✅ Clean, testable code
- ✅ Easy installation
- ✅ Windows/Unix support
- ✅ DPI control
- ✅ Page range

---

**Last updated**: 2025-11-06

See [SUMMARY.md](SUMMARY.md) for a complete project summary.
