package export_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestTxtExporterID(t *testing.T) {
	e := export.TxtExporter{}
	if e.ID() != "txt" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "txt")
	}
}

func TestTxtExporterName(t *testing.T) {
	e := export.TxtExporter{}
	if e.Name() != "Plain Text" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Plain Text")
	}
}

func TestTxtExporterExport(t *testing.T) {
	lines := []string{"line1", "line2", "line3"}
	out, err := export.TxtExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "line1\nline2\nline3"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestTxtExporterExportSingleLine(t *testing.T) {
	out, err := export.TxtExporter{}.Export([]string{"only"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "only" {
		t.Fatalf("unexpected output: got %q, want %q", out, "only")
	}
}

func TestTxtExporterExportEmpty(t *testing.T) {
	out, err := export.TxtExporter{}.Export([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "" {
		t.Fatalf("unexpected output: got %q, want empty string", out)
	}
}
