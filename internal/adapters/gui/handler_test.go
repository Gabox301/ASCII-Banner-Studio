package gui_test

import (
	"context"
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/gui"
	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
)

// newTestHandler ensambla el hexágono real (registries + servicio) detrás del
// adaptador GUI, igual que lo hace la composition root de la app de escritorio.
func newTestHandler() *gui.Handler {
	return gui.New(
		service.NewBannerService(fonts.NewRegistry()),
		export.NewRegistry(),
	)
}

func TestListFontsGuisAvailableFonts(t *testing.T) {
	h := newTestHandler()
	fonts := h.ListFonts()
	if len(fonts) != 4 {
		t.Fatalf("expected 4 fonts, got %d", len(fonts))
	}
	seen := map[string]bool{}
	for _, f := range fonts {
		seen[f.ID] = true
	}
	for _, id := range []string{"block", "big", "banner", "minimal"} {
		if !seen[id] {
			t.Fatalf("expected font %q in the list", id)
		}
	}
}

func TestListExportFormatsIncludesAllFormats(t *testing.T) {
	h := newTestHandler()
	formats := h.ListExportFormats()
	if len(formats) != 6 {
		t.Fatalf("expected 6 export formats, got %d", len(formats))
	}
	seen := map[string]bool{}
	for _, f := range formats {
		seen[f.ID] = true
	}
	for _, id := range []string{"txt", "javascript", "typescript", "rust", "python", "json"} {
		if !seen[id] {
			t.Fatalf("expected format %q in the list", id)
		}
	}
}

func TestGenerateReturnsPlainBanner(t *testing.T) {
	h := newTestHandler()
	out, err := h.Generate("HI", domain.BannerOptions{Font: "block", Align: domain.AlignLeft, Spacing: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "###") || !strings.Contains(out, "#   #") {
		t.Fatalf("unexpected block banner output: %q", out)
	}
}

func TestGenerateReturnsFontNotFoundError(t *testing.T) {
	h := newTestHandler()
	_, err := h.Generate("HI", domain.BannerOptions{Font: "no-such-font"})
	if err != domain.ErrFontNotFound {
		t.Fatalf("expected ErrFontNotFound, got %v", err)
	}
}

func TestExportWithTxtFormatReturnsPlainText(t *testing.T) {
	h := newTestHandler()
	out, err := h.Export("A", domain.BannerOptions{Font: "minimal"}, "txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "A" {
		t.Fatalf("expected plain text output, got %q", out)
	}
}

func TestExportWithJavaScriptFormatWrapsArray(t *testing.T) {
	h := newTestHandler()
	out, err := h.Export("Go", domain.BannerOptions{Font: "block"}, "javascript")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "const BANNER = [") {
		t.Fatalf("expected javascript wrapper, got: %q", out)
	}
	if !strings.Contains(out, "###") {
		t.Fatalf("expected banner rows inside the export, got: %q", out)
	}
}

func TestExportWithUnknownFormatReturnsError(t *testing.T) {
	h := newTestHandler()
	_, err := h.Export("A", domain.BannerOptions{Font: "minimal"}, "nope")
	if err != domain.ErrUnknownExportFormat {
		t.Fatalf("expected ErrUnknownExportFormat, got %v", err)
	}
}

func TestSaveBannerWithoutContextReturnsError(t *testing.T) {
	h := gui.New(
		service.NewBannerService(fonts.NewRegistry()),
		export.NewRegistry(),
	)
	if _, err := h.SaveBanner("content", "banner.txt"); err == nil {
		t.Fatal("expected an error when the Wails context is not set yet")
	}
}

func TestSetContextStoresWailsContext(t *testing.T) {
	h := newTestHandler()
	// Solo debe almacenar el contexto; no abre diálogos en este punto.
	h.SetContext(context.Background())
}

func TestExportWithUnknownFontReturnsError(t *testing.T) {
	h := newTestHandler()
	_, err := h.Export("HI", domain.BannerOptions{Font: "no-such-font"}, "txt")
	if err != domain.ErrFontNotFound {
		t.Fatalf("expected ErrFontNotFound, got %v", err)
	}
}
