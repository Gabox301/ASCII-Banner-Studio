# ASCII Banner Studio (Go, arquitectura hexagonal)

**ASCII Banner Studio** es una plataforma Go construida con arquitectura
hexagonal (ports & adapters) que genera banners ASCII a partir de texto.
El **motor ASCII es independiente de la interfaz**: se expone a través de
dos interfaces — una **CLI** y una **app de escritorio (Wails v2)** — y
cualquier frontend puede conectarse sin tocar el motor.

## Arquitectura hexagonal (ports & adapters)

```
main.go                     → composition root de la app de escritorio (Wails v2)
cmd/asciibanner/            → composition root de la CLI
frontend/dist/              → UI web (HTML/CSS/JS) embebida en el binario
wails.json                  → configuración del build de Wails
internal/
├── core/                   ← el hexágono: cero dependencias externas
│   ├── domain/              Font, Glyph, BannerOptions, Alignment, errores
│   ├── ports/                interfaces: BannerGenerator, FontRepository, Exporter
│   └── service/              BannerService: implementa BannerGenerator
└── adapters/                adaptadores concretos (infraestructura)
    ├── fonts/                adaptador secundario: FontRepository (block/big/banner/minimal)
    ├── export/               adaptador secundario: Exporter (txt/js/ts/python/json)
    ├── cli/                  adaptador primario: traduce argv → BannerGenerator
    └── gui/                  adaptador primario: bindings Wails → BannerGenerator
```

Regla principal: `internal/core` no importa nada de `internal/adapters`
ni de ningún framework. `internal/adapters/cli` y `internal/adapters/gui`
son reemplazables por cualquier otro driving adapter (una futura API
HTTP, una TUI, etc.) sin tocar el motor.

## Aplicación de escritorio (Wails v2)

La app gráfica reutiliza el mismo motor (BannerService + registries) y
ofrece vista previa en vivo mientras escribís, selección de fuente,
espaciado, alineación, mayúsculas/trim y exportación a todos los
formatos con copiado al portapapeles o guardado con diálogo nativo.

### Requisitos

- Go 1.25+
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- Windows: runtime **WebView2** (viene preinstalado en Windows 10/11)
- No requiere Node.js: el frontend es HTML/CSS/JS puro sin build step

### Verificar el entorno

```bash
wails doctor
```

### Modo desarrollo (recarga en vivo de cambios de Go)

```bash
wails dev
```

### Build del instalable de escritorio

```bash
wails build
```

Genera `build/bin/ASCIIBannerStudio.exe`. Para una ventana de release
con ícono y manifiesto, se usan los assets de `build/` (reemplazá
`build/appicon.png` y `build/windows/icon.ico` por tu propio branding).

> Nota: los bindings JS (`frontend/dist/wailsjs/`) se regeneran
> automáticamente con cada `wails dev` / `wails build`. El frontend
> llama al backend a través de `window.go.main.App.*` (Generate,
> Export, ListFonts, ListExportFormats, SaveBanner).

## Uso (CLI)

```bash
go build ./...
go run ./cmd/asciibanner list-fonts
go run ./cmd/asciibanner "SKILLINDEX" -font block -spacing 1 -align left
go run ./cmd/asciibanner "SKILLINDEX" -font banner -format javascript
```

Flags disponibles: `-font` (ver `list-fonts`: 24 fuentes), `-spacing` (int, default 1),
`-align` (left|center|right), `-uppercase`, `-trim`,
`-format` (ver exporters: txt|javascript|typescript|rust|python|json|go|java|csharp|c|kotlin|swift|ruby|php|dart|lua|shell|powershell).

### Ejemplos verificados

Todos los comandos se ejecutan desde la raíz del repo. Salidas copiadas de ejecuciones reales con Go 1.26.5.

#### 1. Listar fuentes disponibles

```bash
go run ./cmd/asciibanner list-fonts
```

```
banner  Banner
big     Big
block   Block
minimal Minimal
```

#### 2. Misma palabra, distintas fuentes

```bash
go run ./cmd/asciibanner "ASCII" -font block
```

```
 ###   ####  #### ##### #####
#   # #     #       #     #
#####  ###  #       #     #
#   #     # #       #     #
#   # ####   #### ##### #####
```

```bash
go run ./cmd/asciibanner "ASCII" -font big
```

```
  ######     ########   ######## ########## ##########
  ######     ########   ######## ########## ##########
##      ## ##         ##             ##         ##
##      ## ##         ##             ##         ##
##########   ######   ##             ##         ##
##########   ######   ##             ##         ##
##      ##         ## ##             ##         ##
##      ##         ## ##             ##         ##
##      ## ########     ######## ########## ##########
##      ## ########     ######## ########## ##########
```

```bash
go run ./cmd/asciibanner "ASCII" -font banner
```

```
  ███     ████    ████   █████   █████
 █   █   █       █         █       █
 █████    ███    █         █       █
 █   █       █   █         █       █
 █   █   ████     ████   █████   █████
```

```bash
go run ./cmd/asciibanner "ASCII" -font minimal
```

```
A S C I I
```

#### 3. Opciones del motor

```bash
# Forzar mayúsculas (útil si la fuente solo tiene A-Z)
go run ./cmd/asciibanner "hola" -font block -uppercase
```

```
#   #  ###  #      ###
#   # #   # #     #   #
##### #   # #     #####
#   # #   # #     #   #
#   #  ###  ##### #   #
```

```bash
# Espaciado entre letras
go run ./cmd/asciibanner "HI" -font block -spacing 3
```

```
#   #   #####
#   #     #
#####     #
#   #     #
#   #   #####
```

```bash
# Multilínea + alineación (left es default)
go run ./cmd/asciibanner "HI`nHI" -font block -align center
go run ./cmd/asciibanner "HI`nHI" -font block -align right
```

Alineación aplicada al ancho máximo del bloque multilínea. Con `center`/`right` el bloque más angosto se rellena con espacios.

#### 4. Exporters (salida lista para copiar a tu código)

```bash
go run ./cmd/asciibanner "Go" -font block -format javascript
```

```javascript
const BANNER = [' ####  ###  ', '#     #   # ', '#  ## #   # ', '#   # #   # ', ' ####  ###  '];
```

```bash
go run ./cmd/asciibanner "Go" -font block -format typescript
```

```typescript
export const BANNER: string[] = [' ####  ###  ', '#     #   # ', '#  ## #   # ', '#   # #   # ', ' ####  ###  '];
```

```bash
go run ./cmd/asciibanner "Go" -font block -format python
```

```python
BANNER = [
    " ####  ###  ",
    "#     #   # ",
    "#  ## #   # ",
    "#   # #   # ",
    " ####  ###  ",
]
```

```bash
go run ./cmd/asciibanner "Go" -font block -format json
```

```json
[" ####  ###  ", "#     #   # ", "#  ## #   # ", "#   # #   # ", " ####  ###  "]
```

> Nota: los exporters escapan correctamente comillas, backslashes, backticks y Unicode. JSON usa `encoding/json` para garantizar sintaxis válida.

#### 5. Combo real

```bash
go run ./cmd/asciibanner "GABO" -font banner -format json
```

```json
[
  "  ████    ███    ████     ███   ",
  " █       █   █   █   █   █   █  ",
  " █  ██   █████   ████    █   █  ",
  " █   █   █   █   █   █   █   █  ",
  "  ████   █   █   ████     ███   "
]
```

```bash
go build -o asciibanner.exe ./cmd/asciibanner
.\asciibanner.exe "SKILLINDEX" -font block -spacing 1 -align left -format javascript
```

## Tests

```bash
go vet ./...
go test ./...
```

Para validar la app de escritorio end-to-end:

```bash
wails build
build/bin/ASCIIBannerStudio.exe
```

Cada fuente vive en su propio archivo bajo `internal/adapters/fonts/` y
se registra en `Registry` sin tocar el algoritmo de `BannerService`.
