@echo off
setlocal enabledelayedexpansion
chcp 65001 >nul

REM Color output
cls
echo.
echo ╔══════════════════════════════════════════════════════════════╗
echo ║           PDF2IMG - Interactive PDF Converter                ║
echo ║                                                              ║
echo ║  Converts PDF pages to PNG or JPG images                    ║
echo ╚══════════════════════════════════════════════════════════════╝
echo.

REM Get PDF information
:input_pdf_path
set "pdf_path="
echo.
echo [1] Enter the path to the PDF file:
echo    Example: C:\documents\file.pdf
echo    Or leave blank to use a PDF in the current folder
set /p pdf_path="PDF path: "

if "%pdf_path%"=="" (
    echo.
    echo Searching for PDFs in the current folder...
    dir /b *.pdf 2>nul
    if !errorlevel! equ 0 (
        echo.
        set /p pdf_path="Enter the PDF name: "
    ) else (
        echo No PDFs found in the current folder.
        goto input_pdf_path
    )
)

REM Validate PDF exists
if not exist "!pdf_path!" (
    echo.
    echo ❌ Error: File does not exist: !pdf_path!
    goto input_pdf_path
)

REM Get PDF info
echo.
echo 📋 Analyzing PDF...
for %%F in ("!pdf_path!") do (
    set "pdf_name=%%~nF"
    set "pdf_size=%%~zF"
)

REM Display file info
echo ✓ File found: !pdf_name!

REM Get page count using pdf2img info command
echo.
echo [2] Getting PDF information...
for /f "tokens=*" %%A in ('pdf2img.exe info "!pdf_path!" 2^>nul ^| findstr /R "^Pages:"') do (
    set "line=%%A"
    for /f "tokens=2" %%B in ("!line!") do set "total_pages=%%B"
)

if "!total_pages!"=="" (
    echo ⚠️  Could not determine the page count automatically
    set /p total_pages="Enter the number of pages in the PDF: "
)

echo ✓ Total pages: !total_pages!

REM Recommend settings based on page count
echo.
echo [3] Automatic recommendations based on page count:
echo.

set "recommended_dpi=150"
set "recommended_format=png"
set "recommended_pool=2"
set "recommended_refresh=50"

if !total_pages! geq 300 (
    set "recommended_dpi=150"
    set "recommended_format=jpg"
    set "recommended_pool=4"
    set "recommended_refresh=25"
    echo 📚 VERY LARGE PDF (300+ pages)
    echo    Recommended DPI: !recommended_dpi! ^(balance^)
    echo    Format: !recommended_format! ^(smaller size^)
    echo    Pool size: !recommended_pool! ^(maximum performance^)
    echo    Refresh: !recommended_refresh! ^(frequent refresh^)
    echo.
) else if !total_pages! geq 200 (
    set "recommended_dpi=150"
    set "recommended_format=jpg"
    set "recommended_pool=3"
    set "recommended_refresh=50"
    echo 📚 Large PDF (200-299 pages)
    echo    Recommended DPI: !recommended_dpi! ^(balance^)
    echo    Format: !recommended_format! ^(smaller size^)
    echo    Pool size: !recommended_pool! ^(good performance^)
    echo    Refresh: !recommended_refresh! ^(every 50 pages^)
    echo.
) else if !total_pages! geq 100 (
    set "recommended_dpi=150"
    set "recommended_format=png"
    set "recommended_pool=2"
    set "recommended_refresh=50"
    echo 📖 Medium PDF (100-199 pages)
    echo    Recommended DPI: !recommended_dpi! ^(balance^)
    echo    Format: !recommended_format! ^(good quality^)
    echo    Pool size: !recommended_pool! ^(standard^)
    echo    Refresh: !recommended_refresh! ^(every 50 pages^)
    echo.
) else (
    set "recommended_dpi=150"
    set "recommended_format=png"
    set "recommended_pool=2"
    set "recommended_refresh=0"
    echo 📄 Small PDF (under 100 pages)
    echo    Recommended DPI: !recommended_dpi! ^(balance^)
    echo    Format: !recommended_format! ^(good quality^)
    echo    Pool size: !recommended_pool! ^(standard^)
    echo    Refresh: disabled ^(not necessary^)
    echo.
)

REM Ask for DPI
echo [4] Select the resolution (DPI):
echo    72-96   = Web/Screen (fast, low quality)
echo    150     = Standard ^(recommended balance^)
echo    200-300 = Print/High quality
echo    600+    = Very detailed (slow)
echo.
set /p user_dpi="DPI [!recommended_dpi!]: "
if "!user_dpi!"=="" set "user_dpi=!recommended_dpi!"

REM Ask for format
echo.
echo [5] Select the format:
echo    png = Lossless, best quality, larger files
echo    jpg = Compressed, smaller files, fast
echo.
set /p user_format="Format [!recommended_format!]: "
if "!user_format!"=="" set "user_format=!recommended_format!"

REM Convert to lowercase
for %%A in (png jpg jpeg) do (
    if /i "!user_format!"=="%%A" set "user_format=%%A"
)

REM Ask for output directory
echo.
echo [6] Where to export the images?
echo    Leave blank to create an "output" folder in the current folder
set /p output_dir="Output folder [.\output]: "
if "!output_dir!"=="" set "output_dir=.\output"

REM Create output directory if doesn't exist
if not exist "!output_dir!" (
    mkdir "!output_dir!"
    echo ✓ Folder created: !output_dir!
)

REM Ask about page range
echo.
echo [7] Convert all pages?
echo    1 = All pages (default)
echo    2 = Specific range
set /p page_choice="Option [1]: "
if "!page_choice!"=="" set "page_choice=1"

set "start_page=0"
set "end_page=0"

if "!page_choice!"=="2" (
    echo.
    set /p start_page="Start page [1]: "
    if "!start_page!"=="" set "start_page=1"

    set /p end_page="End page [!total_pages!]: "
    if "!end_page!"=="" set "end_page=!total_pages!
)

REM Ask for retry on failure
echo.
echo [8] Retry failed pages?
echo    1 = No (default, faster)
echo    2 = Yes (slower, but retries with reduced DPI)
set /p retry_choice="Option [1]: "
if "!retry_choice!"=="" set "retry_choice=1"

set "retry_flag="
if "!retry_choice!"=="2" set "retry_flag=--retry"

REM Summary
echo.
echo.
echo ╔══════════════════════════════════════════════════════════════╗
echo ║                   CONFIGURATION SUMMARY                      ║
echo ╚══════════════════════════════════════════════════════════════╝
echo.
echo  📄 PDF:              !pdf_name!
echo  📄 Pages:            !total_pages!
echo  📊 DPI:              !user_dpi!
echo  🎨 Format:           !user_format!
echo  📁 Output:           !output_dir!

if "!start_page!"=="0" (
    echo  📖 Range:            All pages
) else (
    echo  📖 Range:            Pages !start_page! to !end_page!
)

if "!retry_flag!"=="" (
    echo  🔄 Retry:          No
) else (
    echo  🔄 Retry:          Yes
)

echo.
set /p confirm="Proceed with the conversion? [S/n]: "
if /i "!confirm!"=="n" goto end

REM Build command
set "command=pdf2img.exe -i "!pdf_path!" -o "!output_dir!" -f !user_format! -d !user_dpi!"

if "!total_pages!" geq 300 (
    set "command=!command! --pool-size 4 --refresh-every 25"
) else if "!total_pages!" geq 200 (
    set "command=!command! --pool-size 3 --refresh-every 50"
) else if "!total_pages!" geq 100 (
    set "command=!command! --pool-size 2 --refresh-every 50"
)

if "!start_page!" neq "0" (
    set "command=!command! --start !start_page! --end !end_page!"
)

if not "!retry_flag!"=="" (
    set "command=!command! !retry_flag!
)

REM Execute conversion
echo.
echo ═══════════════════════════════════════════════════════════════
echo 🚀 Starting conversion...
echo ═══════════════════════════════════════════════════════════════
echo.

%command%

REM Check result
if !errorlevel! equ 0 (
    echo.
    echo ═══════════════════════════════════════════════════════════════
    echo ✅ Conversion completed successfully!
    echo ═══════════════════════════════════════════════════════════════
    echo.
    echo 📁 Files saved to: !output_dir!
    echo.
    echo Opening output folder...
    start "" "!output_dir!"
) else (
    echo.
    echo ═══════════════════════════════════════════════════════════════
    echo ❌ Error during conversion
    echo ═══════════════════════════════════════════════════════════════
)

:end
echo.
pause
endlocal
