# Integrating pdf2img with MCP clients

This guide explains how to connect the **pdf2img** MCP server to your MCP client
(for example, Cursor, Zed, or any client compatible with the MCP protocol). Once
configured, your AI assistant can convert PDFs to images and query their
information using natural language.

---

## Requirements

- Go 1.27+ installed (to build the server).
- Your MCP client (for example, Cursor, Zed, or Cline).
- A sample PDF file.

---

## 1. Build the MCP server

From the project root:

```bash
go build -o mcp-server ./cmd/mcp-server
```

Verify that the binary was created:

```bash
ls -la mcp-server
```

> On Windows you can name the binary as you prefer (for example `mcp-server.exe`);
> that absolute path is the one you will use in the client configuration.

---

## 2. Configure your MCP client

Every MCP client reads its MCP configuration file, which contains an
`mcpServers` section. Open your configuration file (the location depends on the
client; most of them open it from the settings menu, under "MCP" or "Developer").

Add an entry for `pdf2img`:

```json
{
     "mcpServers": {
         "pdf2img": {
             "command": "/absolute/path/to/project/mcp-server",
             "args": ["--stdio"]
         }
     }
}
```

**Replace `command`** with the real absolute path of the binary you built:

| System | `command` example |
|---------|----------------------|
| macOS / Linux | `/Users/your-user/projects/mcp-go-pdf-to-img/mcp-server` |
| Windows | `C:\\Users\\your-user\\projects\\mcp-go-pdf-to-img\\mcp-server.exe` |

> If you already have other servers in `mcpServers`, add `pdf2img` as another entry
> inside the same object; do not overwrite the others.

---

## 3. Restart and verify

1. Save the configuration file.
2. Restart (or reload) your MCP client so it detects the new server.
3. Open a conversation and check that `pdf2img` appears as connected, with its two
   tools available:
    - `pdf_to_images`
    - `pdf_info`

If it does not appear, check that the `command` path is correct, that the binary
exists at that path, and that you saved the configuration before restarting.

---

## 4. Available tools

### `pdf_to_images`

Converts pages of a PDF to PNG or JPG images.

| Parameter | Type | Required | Description | Example |
|-----------|------|-----------|-------------|---------|
| `pdf_path` | string | ✅ | Path to the PDF file | `document.pdf` |
| `output_dir` | string | ✅ | Output directory | `./output` |
| `format` | string | ❌ | `png` or `jpg` (default: `png`) | `jpg` |
| `dpi` | number | ❌ | Resolution in DPI (default: `150`) | `300` |
| `start_page` | integer | ❌ | First page (1-indexed) | `1` |
| `end_page` | integer | ❌ | Last page | `10` |
| `prefix` | string | ❌ | File prefix (default: `page_`) | `img_` |

**Example response:**

```json
{
     "total_pages": 25,
     "successful": 25,
     "failed": 0,
     "files": [
         "./output/page_0001.png",
         "./output/page_0002.png",
         "./output/page_0003.png"
     ]
}
```

If any page fails, the response also includes an `errors` field.

### `pdf_info`

Get information about a PDF (number of pages, size, and dimensions).

| Parameter | Type | Required | Description |
|-----------|------|-----------|-------------|
| `pdf_path` | string | ✅ | Path to the PDF file |

**Example response:**

```json
{
     "file": "document.pdf",
     "pages": 25,
     "file_size": "2.50 MB",
     "width": 612.00,
     "height": 792.00
}
```

---

## 5. Usage examples

You can invoke the tools with natural language; your MCP client translates the message
into the appropriate tool call.

**Example 1 — Convert a complete PDF**

> **You:** Convert `document.pdf` to PNG images at 300 DPI and save them to `output`.

> **MCP client →** `pdf_to_images` with `format: "png"`, `dpi: 300`,
> `output_dir: "./output"`.

**Example 2 — Query a PDF**

> **You:** How many pages does `report.pdf` have and what are its dimensions?

> **MCP client →** `pdf_info` with `pdf_path: "report.pdf"`.

**Example 3 — Convert only some pages**

> **You:** Convert pages 1 to 10 of `document.pdf` to JPG at 150 DPI.

> **MCP client →** `pdf_to_images` with `start_page: 1`, `end_page: 10`,
> `format: "jpg"`, `dpi: 150`.

**Example 4 — Generate a thumbnail**

> **You:** Generate a thumbnail (first page) of `document.pdf` in JPG at 96 DPI.

> **MCP client →** `pdf_to_images` with `start_page: 1`, `end_page: 1`,
> `format: "jpg"`, `dpi: 96`, `prefix: "thumbnail_"`.

---

## 6. Troubleshooting

**`pdf2img` does not appear connected**
- Verify that `command` points to the correct absolute path of the binary.
- Confirm that the binary exists at that path (run it from the terminal to check).
- Save the configuration and restart the client.

**The conversion fails or returns an error**
- Use absolute paths for `pdf_path`; relative paths may not resolve.
- The output directory is created automatically if it does not exist; make sure you
  have write permissions in its location.
- If the PDF is corrupted, verify it first with `pdf_info`.

**The images come out blurry**
- Increase the DPI (for example, from `72` to `300`).

**The conversion is slow**
- Reduce the DPI or limit the range with `start_page` / `end_page`.

---

## Related documentation

- [README.md](README.md) — Project overview.
- [EXAMPLES.md](EXAMPLES.md) — Usage examples (CLI and MCP).
- [MCP_TROUBLESHOOTING.md](MCP_TROUBLESHOOTING.md) — Detailed troubleshooting.