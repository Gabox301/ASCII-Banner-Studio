// Package gui es un adaptador primario (driving adapter) para la interfaz de
// escritorio construida con Wails v2. Al igual que internal/adapters/cli, este
// adaptador traduce las peticiones del frontend (bindings JS) en llamadas al
// puerto ports.BannerGenerator, sin tocar jamás el núcleo del hexágono.
package gui

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

// Handler es el objeto que Wails vincula (Bind) al frontend. Cada método
// exportado se convierte en un binding JS bajo window.go.main.App.
//
// El adaptador sigue la misma regla que el CLI: depende únicamente de los
// puertos (ports.BannerGenerator y ports.ExporterRepository), de modo que el
// motor permanece 100% reutilizable entre CLI y GUI.
type Handler struct {
	generator ports.BannerGenerator
	exporters ports.ExporterRepository
	ctx       context.Context
}

// FormatInfo es la proyección liviana de un formato de exportación
// disponible para la GUI (id + nombre legible).
type FormatInfo struct {
	ID   string
	Name string
}

// New crea el handler inyectando los puertos del hexágono.
func New(generator ports.BannerGenerator, exporters ports.ExporterRepository) *Handler {
	return &Handler{generator: generator, exporters: exporters}
}

// SetContext guarda el contexto de la aplicación Wails. Es necesario para
// poder abrir los diálogos nativos (guardado de archivos, mensajes, etc.).
func (h *Handler) SetContext(ctx context.Context) {
	h.ctx = ctx
}

// Generate genera el banner en texto plano (equivalente a la vista previa de
// la GUI y al -format txt de la CLI).
func (h *Handler) Generate(text string, opts domain.BannerOptions) (string, error) {
	return h.generator.Generate(text, opts)
}

// Export genera el banner y lo transforma al formato pedido (txt, javascript,
// typescript, rust, python o json), listo para copiar en código.
func (h *Handler) Export(text string, opts domain.BannerOptions, format string) (string, error) {
	output, err := h.generator.Generate(text, opts)
	if err != nil {
		return "", err
	}
	if format == "" || format == "txt" {
		return output, nil
	}
	exporter, ok := h.exporters.Get(format)
	if !ok {
		return "", domain.ErrUnknownExportFormat
	}
	return exporter.Export(strings.Split(output, "\n"))
}

// ListFonts devuelve las fuentes disponibles (equivalente a list_fonts()).
func (h *Handler) ListFonts() []domain.FontInfo {
	return h.generator.ListFonts()
}

// ListExportFormats devuelve los formatos de exportación disponibles, en el
// orden registrado por el adaptador de exporters.
func (h *Handler) ListExportFormats() []FormatInfo {
	list := h.exporters.List()
	infos := make([]FormatInfo, 0, len(list))
	for _, e := range list {
		infos = append(infos, FormatInfo{ID: e.ID(), Name: e.Name()})
	}
	return infos
}

// defaultFileName elige el nombre sugerido por el usuario o un nombre por
// defecto cuando no se indicó ninguno.
func defaultFileName(suggestedName string) string {
	if suggestedName == "" {
		return "banner.txt"
	}
	return suggestedName
}

// SaveBanner abre el diálogo nativo de guardado y escribe el contenido en el
// archivo elegido por el usuario. Devuelve la ruta del archivo guardado, o una
// cadena vacía cuando el usuario cancela el diálogo.
func (h *Handler) SaveBanner(content string, suggestedName string) (string, error) {
	if h.ctx == nil {
		return "", errors.New("la aplicación todavía no está lista; reintentá en un momento")
	}
	file, err := runtime.SaveFileDialog(h.ctx, runtime.SaveDialogOptions{
		Title:           "Guardar banner ASCII",
		DefaultFilename: defaultFileName(suggestedName),
	})
	if err != nil {
		return "", err
	}
	if file == "" {
		return "", nil // usuario canceló el diálogo
	}
	if err := os.WriteFile(file, []byte(content), 0o644); err != nil {
		return "", err
	}
	return file, nil
}
