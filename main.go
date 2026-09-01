// App de escritorio ASCII Banner Studio (Wails v2).
//
// Este archivo es la composition root de la interfaz gráfica: ensambla el
// hexágono (fonts + export + service) y lo expone al frontend web a través de
// Wails. Es el equivalente gráfico de cmd/asciibanner (la CLI sigue intacta).
//
//   - internal/core       → el motor (cero dependencias de framework)
//   - internal/adapters   → font/export registries + adaptador gui
//   - frontend/dist       → assets embebidos (HTML/CSS/JS)
package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/gui"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
)

//go:embed all:frontend/dist
var uiAssets embed.FS

//go:embed all:build/appicon.png
var appIconAssets embed.FS

// assetFS combina los assets del frontend (frontend/dist) con el ícono de la
// app (build/appicon.png), expuesto en la raíz como logo.png. Así el header de
// la UI muestra exactamente el mismo archivo que se usa como ícono de la app:
// una única fuente de verdad, sin archivos PNG duplicados en el repo.
type assetFS struct {
	ui   fs.FS // frontend/dist embebido
	icon fs.FS // build/appicon.png embebido
}

func newAssetFS() *assetFS {
	// `//go:embed all:frontend/dist` guarda los archivos bajo el prefijo
	// "frontend/dist/..."; lo recortamos para servirlos desde la raíz.
	ui, err := fs.Sub(uiAssets, "frontend/dist")
	if err != nil {
		panic(fmt.Sprintf("assets del frontend no encontrados: %v", err))
	}
	return &assetFS{ui: ui, icon: appIconAssets}
}

func normalizePath(name string) string {
	return path.Clean(strings.TrimPrefix(name, "/"))
}

// renamedFile adapta un fs.FileInfo para exponerlo con otro nombre dentro del
// directorio virtual (logo.png ← build/appicon.png).
type renamedFile struct {
	fs.FileInfo
	name string
}

func (r renamedFile) Name() string { return r.name }

// Open implementa fs.FS.
func (a *assetFS) Open(name string) (fs.File, error) {
	key := normalizePath(name)
	if key == "logo.png" {
		return a.icon.Open("build/appicon.png")
	}
	return a.ui.Open(key)
}

// Stat implementa fs.StatFS.
func (a *assetFS) Stat(name string) (fs.FileInfo, error) {
	key := normalizePath(name)
	if key == "logo.png" {
		return fs.Stat(a.icon, "build/appicon.png")
	}
	return fs.Stat(a.ui, key)
}

// ReadFile implementa fs.ReadFileFS.
func (a *assetFS) ReadFile(name string) ([]byte, error) {
	key := normalizePath(name)
	if key == "logo.png" {
		return fs.ReadFile(a.icon, "build/appicon.png")
	}
	return fs.ReadFile(a.ui, key)
}

// ReadDir implementa fs.ReadDirFS. En la raíz suma el logo.png de los
// assets del build a la lista del frontend.
func (a *assetFS) ReadDir(name string) ([]fs.DirEntry, error) {
	key := normalizePath(name)
	if key == "logo.png" {
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: fs.ErrInvalid}
	}
	entries, err := fs.ReadDir(a.ui, key)
	if err != nil {
		return nil, err
	}
	if key == "." {
		info, err := fs.Stat(a.icon, "build/appicon.png")
		if err != nil {
			return nil, err
		}
		entries = append(entries, fs.FileInfoToDirEntry(renamedFile{FileInfo: info, name: "logo.png"}))
	}
	return entries, nil
}

// App es la cáscara que Wails vincula al frontend. Vive en el package main
// para que los bindings JS se generen bajo un path limpio
// (window.go.main.App.*). Toda la lógica delega en el adaptador gui.Handler,
// preservando la arquitectura hexagonal del proyecto.
type App struct {
	*gui.Handler
}

// startup captura el contexto de Wails para habilitar los diálogos nativos
// (guardado de archivos). Se invoca una sola vez al iniciar la app.
func (a *App) startup(ctx context.Context) {
	a.SetContext(ctx)
}

// buildApp ensambla la composition root de la app de escritorio: mismos
// registries y servicio que usa la CLI. Separada de main() para poder
// testear el cableado real en proceso sin abrir una ventana Wails.
func buildApp() *App {
	fontRepo := fonts.NewRegistry()
	exportRepo := export.NewRegistry()
	generator := service.NewBannerService(fontRepo)
	return &App{Handler: gui.New(generator, exportRepo)}
}

func main() {
	app := buildApp()

	err := wails.Run(&options.App{
		Title:     "ASCII Banner Studio",
		Width:     1120,
		Height:    740,
		MinWidth:  860,
		MinHeight: 560,
		AssetServer: &assetserver.Options{
			Assets: newAssetFS(),
		},
		BackgroundColour: &options.RGBA{R: 15, G: 17, B: 23, A: 255},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatalf("error al iniciar la aplicación: %v", err)
	}
}
