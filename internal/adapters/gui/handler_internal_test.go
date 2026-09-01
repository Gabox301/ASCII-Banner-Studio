package gui

import "testing"

// TestDefaultFileName cubre la lógica de nombre por defecto del diálogo de
// guardado. Se mantiene en este paquete (test interno) porque defaultFileName
// es un detalle de implementación no exportado.
func TestDefaultFileName(t *testing.T) {
	if got := defaultFileName("banner-2026.txt"); got != "banner-2026.txt" {
		t.Fatalf("unexpected filename: %q", got)
	}
	if got := defaultFileName(""); got != "banner.txt" {
		t.Fatalf("expected default filename, got %q", got)
	}
}
