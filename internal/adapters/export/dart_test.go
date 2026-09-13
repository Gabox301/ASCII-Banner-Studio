package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestDartExporterIDAndName(t *testing.T) {
	e := export.DartExporter{}
	if e.ID() != "dart" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "dart")
	}
	if e.Name() != "Dart" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Dart")
	}
}

func TestDartExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.DartExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "const BANNER = [\n  '##',\n  '  ',\n];\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestDartExporterEscapesDollarInterpolation(t *testing.T) {
	lines := []string{`$x and ${y}`}
	out, err := export.DartExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `\$x and \${y}`) {
		t.Fatalf("expected escaped dollar interpolation, got: %s", out)
	}
}
