package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestPythonExporterID(t *testing.T) {
	e := export.PythonExporter{}
	if e.ID() != "python" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "python")
	}
}

func TestPythonExporterName(t *testing.T) {
	e := export.PythonExporter{}
	if e.Name() != "Python" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Python")
	}
}

func TestPythonExporterMirrorBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.PythonExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "BANNER = [\n    \"##\",\n    \"  \",\n]\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestPythonExporterEscapesDoubleQuotes(t *testing.T) {
	lines := []string{`say "hi"`}
	out, err := export.PythonExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `say \"hi\"`) {
		t.Fatalf("expected escaped double quotes, got: %s", out)
	}
}
