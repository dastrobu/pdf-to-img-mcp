# Quick Start Guide

## 5 minutos para tener pdf2img funcionando

### Opción 1: Linux / macOS

```bash
# 1. Clona o descarga el proyecto
cd pdf2img

# 2. Ejecuta el script de instalación
chmod +x install.sh
./install.sh

# 3. ¡Listo! Usa:
./pdf2img -i tu-documento.pdf -o ./output
```

### Opción 2: Windows

```bash
# 1. Abre PowerShell como administrador
# 2. Navega a la carpeta del proyecto
cd pdf2img

# 3. Ejecuta el script de instalación
install.bat

# 4. ¡Listo! Usa:
pdf2img.exe -i tu-documento.pdf -o .\output
```

### Opción 3: Go directo (cualquier SO)

```bash
# Requiere Go 1.21+
go mod download
go build -o pdf2img .

# Usar:
./pdf2img -i documento.pdf -o ./output
```

## Primeros pasos

### Convertir un PDF completo

```bash
pdf2img -i documento.pdf -o ./salida
```

**Resultado:**
- ✓ Crea directorio `./salida`
- ✓ Genera `page_0001.png`, `page_0002.png`, etc.

### Obtener información del PDF

```bash
pdf2img info documento.pdf
```

**Resultado:**
```
PDF Information
===============
File: documento.pdf
Pages: 42
Size: 2.50 MB
```

### Convertir solo las primeras 5 páginas

```bash
pdf2img -i documento.pdf -o ./salida --start 1 --end 5
```

### Usar JPG en lugar de PNG

```bash
pdf2img -i documento.pdf -o ./salida -f jpg
```

### Aumentar calidad (más DPI)

```bash
pdf2img -i documento.pdf -o ./salida -d 300
```

| DPI | Uso |
|-----|-----|
| 72-96 | Web, pantalla |
| 150 | Balance calidad/velocidad (default) |
| 300 | Imprenta, documentos |
| 600+ | Archivos muy detallados |

## Referencia completa de opciones

### Todas las banderas disponibles

| Flag | Corto | Descripción | Defecto | Ejemplo |
|------|-------|-------------|---------|---------|
| `--input` | `-i` | Archivo PDF (requerido) | - | `-i documento.pdf` |
| `--output` | `-o` | Directorio de salida | `.` (actual) | `-o ./output` |
| `--format` | `-f` | Formato: `png` o `jpg` | `png` | `-f jpg` |
| `--dpi` | `-d` | DPI para renderizado | `150` | `-d 300` |
| `--start` | - | Página inicial (1-indexada) | `0` (primera) | `--start 1` |
| `--end` | - | Página final (1-indexada) | `0` (última) | `--end 10` |
| `--prefix` | - | Prefijo para archivos | `page_` | `--prefix img_` |
| `--verbose` | `-v` | Salida detallada | `false` | `-v` |
| `--retry` | - | Reintentar páginas fallidas con DPI reducido | `false` | `--retry` |
| `--pool-size` | - | Max instancias PDFium en pool (para PDFs grandes) | `2` | `--pool-size 4` |
| `--refresh-every` | - | Refrescar instancia WASM cada N páginas (0=desactivar) | `50` | `--refresh-every 25` |

### Casos de uso comunes

#### Miniatura de primera página
```bash
pdf2img -i documento.pdf -o ./thumbs --start 1 --end 1 -d 96 --prefix thumb_
```

#### Crear galería web
```bash
pdf2img -i documento.pdf -o ./web/images -f jpg -d 150
```

#### Extraer rango específico
```bash
pdf2img -i libro.pdf -o ./capitulo5 --start 100 --end 150
```

## Combinaciones de opciones para diferentes escenarios

### 📱 Para Web/Pantalla (baja calidad, rápido)
```bash
pdf2img -i documento.pdf -o ./output -f jpg -d 96
```
- Formato: JPG (más pequeño)
- DPI: 96 (suficiente para pantalla)
- Rápido y files pequeños

### 📄 Documentos estándar (balance)
```bash
pdf2img -i documento.pdf -o ./output -f png -d 150
```
- Formato: PNG (lossless)
- DPI: 150 (default, balance calidad/velocidad)
- Bueno para la mayoría de casos

### 🖨️ Para imprenta/alta calidad
```bash
pdf2img -i documento.pdf -o ./output -f png -d 300
```
- Formato: PNG (sin pérdida)
- DPI: 300 (alta calidad)
- Más lento pero excelente calidad

### 📚 PDFs muy grandes (>100 páginas) - Optimizado
```bash
pdf2img -i documento_grande.pdf -o ./output -d 150 --refresh-every 50 --pool-size 3
```
- `--refresh-every 50`: Refresca WASM cada 50 páginas (limpia memoria)
- `--pool-size 3`: Más instancias PDFium para mejor rendimiento
- Evita corrupción de rendering después de muchas páginas

### 📚 PDFs muy grandes (>200 páginas) - Máximo control
```bash
pdf2img -i documento_enorme.pdf -o ./output -d 150 --refresh-every 25 --pool-size 4 --retry
```
- `--refresh-every 25`: Refresca más frecuentemente (más memoria)
- `--pool-size 4`: Máximas instancias para mejor rendimiento
- `--retry`: Reintenta páginas fallidas con DPI reducido

### ⚡ Procesamiento rápido (baja calidad)
```bash
pdf2img -i documento.pdf -o ./output -d 72
```
- DPI: 72 (mínimo, muy rápido)
- Útil para previsualizaciones rápidas

### 🎯 Paginas específicas (rango)
```bash
pdf2img -i documento.pdf -o ./output --start 50 --end 150
```
- Solo convierte páginas 50 a 150
- Útil para procesar documentos muy grandes por partes

### 🔍 Paginas específicas + alta calidad
```bash
pdf2img -i documento.pdf -o ./output --start 1 --end 10 -d 300 -f png
```
- Primeras 10 páginas con máxima calidad
- PNG lossless para preservar detalles

### 📦 Miniaturas
```bash
pdf2img -i documento.pdf -o ./thumbs --start 1 --end 1 -d 96 --prefix thumb_ -f jpg
```
- Solo primera página
- 96 DPI (bueno para thumbs)
- JPG (más pequeño)
- Prefijo: `thumb_`

### 🎬 Galería de imágenes
```bash
pdf2img -i documento.pdf -o ./gallery -f jpg -d 150 --prefix gallery_
```
- Todas las páginas
- JPG (menor tamaño)
- 150 DPI (balance)
- Prefijo personalizado

### 📊 Reportes/Análisis
```bash
pdf2img -i reporte.pdf -o ./análisis -d 200 -f png --prefix report_
```
- Más alto que default (200 vs 150 DPI)
- PNG para preservar calidad
- Mejor para análisis detallados

### 🖼️ Archivos detallados (gráficos, planos)
```bash
pdf2img -i plano.pdf -o ./output -d 600 -f png
```
- 600 DPI (muy alta calidad)
- PNG (sin pérdida)
- Más lento, pero excelente para detalles

### 🧪 Debugging/Diagnostico
```bash
pdf2img -i documento.pdf -o ./output -v --retry
```
- `-v`: Salida detallada (verbose)
- `--retry`: Reintentar páginas fallidas
- Útil para diagnosticar problemas

### 📋 Información del PDF (sin conversión)
```bash
pdf2img info documento.pdf
```
- Solo muestra: páginas, tamaño, dimensiones
- No convierte nada

## Combinaciones avanzadas para PDFs problemáticos

### Documento con muchas páginas y posibles errores
```bash
pdf2img -i documento.pdf -o ./output -d 150 --refresh-every 25 --pool-size 4 --retry -v
```
Combina:
- `--refresh-every 25`: Refresca WASM frecuentemente
- `--pool-size 4`: Más recursos
- `--retry`: Reintenta con DPI reducido
- `-v`: Ve qué pasa en detalle

### Procesar en bloques (para PDFs de 300+ páginas)
```bash
# Bloque 1
pdf2img -i documento.pdf -o ./output --start 1 --end 75 -d 150 --refresh-every 50

# Bloque 2
pdf2img -i documento.pdf -o ./output --start 76 --end 150 -d 150 --refresh-every 50

# Bloque 3
pdf2img -i documento.pdf -o ./output --start 151 --end 225 -d 150 --refresh-every 50

# Bloque 4
pdf2img -i documento.pdf -o ./output --start 226 --end 300 -d 150 --refresh-every 50
```
Procesar en bloques reduce problemas de memoria con PDFs muy grandes.

### PDF grande con calidad variable
```bash
# Páginas normales con DPI standard
pdf2img -i documento.pdf -o ./output --start 1 --end 50 -d 150

# Páginas con gráficos con mayor DPI
pdf2img -i documento.pdf -o ./output --start 51 --end 100 -d 300
```

## Consejos de rendimiento

### Si es lento:
```bash
# Reducir DPI
pdf2img -i documento.pdf -o ./output -d 96

# O especificar solo rango necesario
pdf2img -i documento.pdf -o ./output --start 1 --end 10
```

### Si fallan páginas:
```bash
# Usar retry automático
pdf2img -i documento.pdf -o ./output --retry

# O aumentar refresh
pdf2img -i documento.pdf -o ./output --refresh-every 25
```

### Si usa mucha memoria:
```bash
# Reducir pool size y aumentar refresh
pdf2img -i documento.pdf -o ./output --pool-size 1 --refresh-every 25
```

## Usar como servidor MCP

### Opción HTTP (puerto 8080)

```bash
mcp-server
# Servidor en: http://localhost:8080

# En otra terminal:
curl http://localhost:8080/tools
```

### Opción stdio (integración con Claude)

```bash
mcp-server --stdio
```

## Documentación completa

- [README.md](README.md) - Documentación completa
- [EXAMPLES.md](EXAMPLES.md) - Ejemplos detallados
- [DEVELOPMENT.md](DEVELOPMENT.md) - Guía para desarrolladores

## Problemas comunes

| Problema | Solución |
|----------|----------|
| `command not found: pdf2img` | Asegúrate de estar en el directorio del proyecto o instalar globalmente con `--global` |
| PDFium no se descarga | Ejecuta `go mod download` |
| Las imágenes se ven borrosas | Aumenta DPI: `-d 300` |
| El proceso es lento | Reduce DPI: `-d 96` o especifica un rango: `--start 1 --end 10` |
| `file not found` | Verifica que el PDF existe: `ls documento.pdf` |

## Próximos pasos

1. ✅ **Primeros pasos**: Convertir tu primer PDF
2. 📖 **Aprender más**: Lee [EXAMPLES.md](EXAMPLES.md)
3. 🔧 **Personalizar**: Ajusta DPI, formato y opciones
4. 🌐 **Integrar**: Usa como servidor MCP o en tu código Go
5. 🚀 **Contribuir**: ¡Las PRs son bienvenidas!

## Ayuda

```bash
pdf2img --help
pdf2img info --help
mcp-server --help
```

---

**¿Preguntas?** Ver [README.md](README.md) para documentación completa.
