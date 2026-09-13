package main

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

func TestAssetFS_ServesLogoFromAppIcon(t *testing.T) {
	a := newAssetFS()

	// El logo que sirve la UI es byte a byte el appicon.png del build:
	// una sola fuente de verdad, sin copias duplicadas.
	got, err := fs.ReadFile(a, "logo.png")
	if err != nil {
		t.Fatalf("leer logo.png: %v", err)
	}
	want, err := fs.ReadFile(appIconAssets, "build/appicon.png")
	if err != nil {
		t.Fatalf("leer build/appicon.png: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal("logo.png no coincide con build/appicon.png")
	}

	// Debe ser un PNG válido.
	if len(got) < 8 || got[0] != 0x89 || string(got[1:4]) != "PNG" {
		t.Fatalf("logo.png no parece un PNG: %d bytes", len(got))
	}
}

func TestAssetFS_ServesFrontendAssets(t *testing.T) {
	a := newAssetFS()
	for _, f := range []string{"index.html", "style.css", "app.js"} {
		if _, err := fs.Stat(a, f); err != nil {
			t.Fatalf("%s debería servirse desde frontend/dist: %v", f, err)
		}
	}
}

func TestAssetFS_ReadDirCombinesFrontendAndLogo(t *testing.T) {
	a := newAssetFS()
	entries, err := fs.ReadDir(a, ".")
	if err != nil {
		t.Fatalf("readdir de la raíz: %v", err)
	}
	found := false
	for _, e := range entries {
		if e.Name() == "logo.png" {
			found = true
		}
	}
	if !found {
		t.Fatal("logo.png debería aparecer en el directorio raíz")
	}
}

func TestAssetFS_OpenLogo(t *testing.T) {
	a := newAssetFS()
	f, err := a.Open("logo.png")
	if err != nil {
		t.Fatalf("open logo.png: %v", err)
	}
	defer f.Close()
	var head [4]byte
	if _, err := io.ReadFull(f, head[:]); err != nil {
		t.Fatalf("leer cabecera de logo.png: %v", err)
	}
	if !bytes.Equal(head[:], []byte{0x89, 'P', 'N', 'G'}) {
		t.Fatalf("logo.png no es un PNG válido: % x", head)
	}
}

func TestAssetFS_OpenFrontendAsset(t *testing.T) {
	a := newAssetFS()
	f, err := a.Open("app.js")
	if err != nil {
		t.Fatalf("open app.js: %v", err)
	}
	f.Close()
}

func TestAssetFS_OpenMissingReturnsError(t *testing.T) {
	a := newAssetFS()
	if _, err := a.Open("no-such.txt"); err == nil {
		t.Fatal("se esperaba error al abrir un archivo inexistente")
	}
}

func TestAssetFS_ReadFileFrontendAsset(t *testing.T) {
	a := newAssetFS()
	data, err := fs.ReadFile(a, "style.css")
	if err != nil {
		t.Fatalf("read style.css: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("style.css no debería estar vacío")
	}
}

func TestAssetFS_StatLogo(t *testing.T) {
	a := newAssetFS()
	info, err := fs.Stat(a, "logo.png")
	if err != nil {
		t.Fatalf("stat logo.png: %v", err)
	}
	if info.Size() <= 0 {
		t.Fatalf("se esperaba un logo no vacío, got %d bytes", info.Size())
	}
}

func TestAssetFS_ReadDirOnLogoReturnsError(t *testing.T) {
	a := newAssetFS()
	if _, err := fs.ReadDir(a, "logo.png"); err == nil {
		t.Fatal("se esperaba error al listar un archivo como directorio")
	}
}

func TestAssetFS_ReadDirOnMissingDirReturnsError(t *testing.T) {
	a := newAssetFS()
	if _, err := fs.ReadDir(a, "no-such-dir"); err == nil {
		t.Fatal("se esperaba error al listar un directorio inexistente")
	}
}

func TestBuildAppWiring(t *testing.T) {
	app := buildApp()
	if app == nil || app.Handler == nil {
		t.Fatal("buildApp devolvió un App sin handler")
	}
	if len(app.ListFonts()) != 24 {
		t.Fatalf("expected 24 fonts, got %d", len(app.ListFonts()))
	}
	out, err := app.Generate("HI", domain.BannerOptions{Font: "minimal", Spacing: 0})
	if err != nil {
		t.Fatalf("generate via buildApp: %v", err)
	}
	if !strings.Contains(out, "HI") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestStartupSetsContext(t *testing.T) {
	app := buildApp()
	// Solo almacena el contexto de Wails; no abre ventanas ni diálogos.
	app.startup(context.Background())
}

type errorFS struct{}

func (errorFS) Open(name string) (fs.File, error) {
	return nil, fs.ErrNotExist
}

func TestNormalizePath(t *testing.T) {
	if got := normalizePath("/index.html"); got != "index.html" {
		t.Fatalf("unexpected normalizePath: %q", got)
	}
	if got := normalizePath("index.html"); got != "index.html" {
		t.Fatalf("unexpected normalizePath: %q", got)
	}
	if got := normalizePath("/frontend/dist/app.js"); got != "frontend/dist/app.js" {
		t.Fatalf("unexpected normalizePath: %q", got)
	}
}

func TestRenamedFileName(t *testing.T) {
	r := renamedFile{name: "logo.png"}
	if r.Name() != "logo.png" {
		t.Fatalf("renamedFile.Name: got %q", r.Name())
	}
}

func TestAssetFS_ReadDirRootWithMissingIcon(t *testing.T) {
	ui, err := fs.Sub(uiAssets, "frontend/dist")
	if err != nil {
		t.Fatalf("setup ui fs: %v", err)
	}
	a := &assetFS{ui: ui, icon: errorFS{}}
	if _, err := fs.ReadDir(a, "."); err == nil {
		t.Fatal("expected error when icon is missing")
	}
}
