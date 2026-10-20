# 🔧 Troubleshooting - MCP Server with your MCP client

Guide to solving common problems when you configure pdf2img with your MCP client.

---

## ✅ Pre-check (Before You Continue)

Before troubleshooting, verify these points:

### 1. MCP Server compiled
```bash
# From the project root
ls -la mcp-server.exe
# Should be ~18 MB
```

### 2. Correct configuration file
```bash
# Open your client's MCP configuration file (the path varies by client)

# Verify that you see something like:
{
    "mcpServers": {
        "pdf2img": {
            "command": "C:\\ruta\\a\\mcp-server.exe",
            "args": ["--stdio"]
        }
    }
}
```

### 3. Correct path
- ✅ Use the **FULL** path to `mcp-server.exe`
- ✅ Replace `C:\ruta\a` with your real path
- ❌ DO NOT use relative paths like `./mcp-server.exe`
- ❌ DO NOT use `~` (tilde) in Windows

### 4. File saved
After editing your client's MCP configuration file:
- Press Ctrl+S
- Close the editor
- Verify that the file has the changes

---

## 🐛 Common Problems and Solutions

### Problem 1: "pdf2img does not appear connected"

**Symptoms**:
- You don't see "pdf2img" in your MCP client
- Or you see "pdf2img" but with status "disconnected"

**Solutions** (in order):

1. **Restart your MCP client completely**
    ```bash
    # Close all windows of your MCP client
    # Wait 2 seconds
    # Open your MCP client again
    ```

2. **Verify the configuration again**
    ```json
    // ✅ Correct
    {
        "command": "C:\\Users\\DAVID\\Projects\\pdf2img\\mcp-server.exe",
        "args": ["--stdio"]
    }

    // ❌ Incorrect
    {
        "command": "mcp-server.exe",   // Missing full path
        "args": ["--stdio"]
    }
    ```

3. **Test the server manually**
    ```bash
    cd C:\ruta\a\tu\proyecto
    mcp-server.exe --stdio
    # Should say: "MCP Server running in stdio mode"
    # Press Ctrl+C to exit
    ```

4. **Look for errors in the log**
    - In your MCP client, open the configuration/developer section (varies by client)
    - Look for error messages in red
    - Copy the exact error messages

### Problem 2: "MCP Server connection failed"

**Symptoms**:
- You see a connection error
- The server does not respond to your MCP client

**Solutions**:

1. **Verify that the file exists**
    ```bash
    dir "C:\ruta\exacta\mcp-server.exe"
    # If it does not exist, compile again:
    cd C:\tu\proyecto
    go build -o mcp-server.exe ./cmd/mcp-server
    ```

2. **Verify permissions**
    - Right-click on `mcp-server.exe`
    - Properties → General
    - Is there an "Unlock" button?
    - If so, click Unlock
    - Click OK
    - Restart your MCP client

3. **Verify that Go is installed**
    ```bash
    go version
    # Should show something like: go version go1.21.5 windows/amd64
    ```

### Problem 3: "JSON-RPC errors" in the logs

**Symptoms**:
- In the logs you see "invalid_literal", "Unrecognized key(s)" errors
- The server responds but your MCP client does not accept the responses

**✅ FIXED IN v2.0.0**:
The server now uses `github.com/modelcontextprotocol/go-sdk` (official MCP Go SDK) that correctly implements JSON-RPC 2.0.

**What changed**:
```
❌ Before: Manual JSON-RPC 2.0 implementation
✅ Now: Official SDK that handles the whole protocol automatically
```

**To get the fixed version**:
```bash
cd C:\tu\proyecto
go build -o mcp-server.exe ./cmd/mcp-server
# Open your MCP client again
```

**If you still see errors**:
- Verify that you compiled the latest version (recent timestamp)
- Run `go mod tidy` if there are compilation problems
- Report the exact log error with full details

### Problem 4: "Tool execution error" in your MCP client

**Symptoms**:
- Your MCP client connects to the server
- You see the available tools
- But when using a tool, an error appears

**Causes and solutions**:

1. **Incorrect PDF path**
    ```
    ❌ "document.pdf"   (relative path)
    ✅ "C:\Users\DAVID\Documents\document.pdf"   (absolute path)
    ```

2. **Output directory does not exist**
    - The server should create it automatically
    - But verify that you have write permissions in that directory

3. **Corrupt PDF file**
    ```bash
    # Try with the example PDF
    # From your AI assistant: "How many pages does example.pdf have?"
    ```

4. **DPI too high**
    - If you use DPI > 300, there may be memory problems
    - Try a lower DPI (150 or 200)

### Problem 5: "File not found"

**Symptoms**:
- Error: "input file not found"
- Or: "PDF file not found"

**Solutions**:

1. **Verify that the file exists**
    ```bash
    ls "C:\Users\DAVID\Documents\document.pdf"
    ```

2. **Use absolute path**
    - Ask your AI assistant: "Use the full path: C:\Users\DAVID\Documents\document.pdf"

3. **Verify permissions**
    - The file must be readable
    - It must not be open in another program

### Problem 6: "Timeout" or "Server not responding"

**Symptoms**:
- Your AI assistant starts the operation but never responds
- After waiting, a timeout occurs

**Solutions**:

1. **Verify that the server is running**
    ```bash
    tasklist | findstr mcp-server
    # Should show: mcp-server.exe ...
    ```

2. **Restart your MCP client**
     ```bash
     # Wait 3 seconds
     # Open your MCP client again
     ```

3. **Verify that the PDF is not too large**
    - PDFs > 50 MB may take time
    - Try a smaller PDF

---

## 📊 Diagnostic Checklist

If nothing works, go through this checklist:

- [ ] Did I compile `mcp-server.exe` correctly?
   ```bash
    go build -o mcp-server.exe ./cmd/mcp-server
   ```

- [ ] Does the `mcp-server.exe` file exist?
   ```bash
    ls -lh mcp-server.exe
   ```

- [ ] Is the configuration in your MCP client correct?
   - Full path (not relative)
   - No special characters
   - `"args": ["--stdio"]` exact

- [ ] Did I restart my MCP client after changing config?
     ```bash
     # Close and reopen your MCP client
     ```

- [ ] Does the server work manually?
   ```bash
    mcp-server.exe --stdio
   ```

- [ ] Do I see "pdf2img" in my MCP client?
   - Check the list of MCP servers
   - Or the configuration/developer section

- [ ] Can I list tools in my AI assistant?
   - Ask: "What tools do I have available?"
   - Should mention `pdf_info` and `pdf_to_images`

- [ ] Can I use `pdf_info`?
   - Ask: "How many pages does example.pdf have?"
   - Should return: "1 page"

---

## 🆘 When Nothing Works

If you have already tried everything, here is what to report:

1. **Open a terminal**
    ```bash
    cd C:\tu\proyecto
    mcp-server.exe --stdio
    ```

2. **Take a screenshot of:**
    - The server output
    - Your MCP client's log

3. **Include information about:**
    - Operating system: `ver` (on Windows)
    - Go version: `go version`
    - Project path: `cd` (show in terminal)

4. **Create an issue on GitHub**
    - Include everything above
    - Describe what you tried
    - Include exact error messages

---

## 📝 Useful Logs

### View server logs in real time
```bash
# Terminal 1: Start the server
mcp-server.exe --stdio

# Terminal 2: Try using from your AI assistant
# (See the logs in Terminal 1)
```

### Save logs for analysis
```bash
# Redirect stderr to a file
mcp-server.exe --stdio 2> mcp-errors.log

# Then open the file
cat mcp-errors.log
```

---

## ✅ Successful Verification

When it works, you will see:

1. **In your MCP client**
    - ✓ "pdf2img" appears connected
    - ✓ No red errors in the logs

2. **When using a tool**
    - ✓ Your AI assistant starts the operation
    - ✓ The tool executes
    - ✓ You see the result in less than 10 seconds

3. **In the server log**
    ```
    ✓ Server started and connected successfully
    ✓ Received method: initialize
    ✓ Received method: tools/list
    ✓ Received method: tools/call
    ```

---

**Last updated**: 2025-11-07
**Version**: 1.0.0