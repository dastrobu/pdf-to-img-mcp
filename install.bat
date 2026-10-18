@echo off
REM Installation script for pdf2img on Windows

echo 📦 Installing pdf2img...

REM Check for Go
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Go is not installed. Please install Go 1.21 or higher.
    echo    Visit: https://golang.org/dl
    exit /b 1
)

for /f "tokens=3" %%a in ('go version') do set GO_VERSION=%%a
echo ✓ Found Go version: %GO_VERSION%

REM Download dependencies
echo 📥 Downloading dependencies...
call go mod download
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Failed to download dependencies
    exit /b 1
)
call go mod verify

REM Build CLI
echo 🔨 Building CLI...
call go build -o pdf2img.exe .
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Failed to build CLI
    exit /b 1
)
echo ✓ CLI built: pdf2img.exe

REM Build MCP server (optional)
if "%1"=="--with-mcp" (
    echo 🔨 Building MCP server...
    call go build -o mcp-server.exe .\cmd\mcp-server
    if %ERRORLEVEL% NEQ 0 (
        echo ❌ Failed to build MCP server
        exit /b 1
    )
    echo ✓ MCP server built: mcp-server.exe
)

REM Test installation
echo 🧪 Testing installation...
pdf2img.exe --help >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo ⚠️  Warning: CLI test failed
) else (
    echo ✓ CLI works correctly
)

REM Optional: Install globally
if "%1"=="--global" (
    echo 🌐 Installing globally...
    call go install .
    echo ✓ Installed globally as 'mcp-go-pdf-to-img' (build with an explicit name via: go build -o pdf2img .)

    if "%2"=="--with-mcp" (
        call go install .\cmd\mcp-server
        echo ✓ Installed globally as 'mcp-server'
    )
)

echo.
echo ✅ Installation complete!
echo.
echo Usage:
echo   pdf2img.exe -i input.pdf -o .\output
echo   pdf2img.exe info input.pdf
echo.
echo For more information, see README.md
