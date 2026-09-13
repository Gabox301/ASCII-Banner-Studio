package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestPHPExporterIDAndName(t *testing.T) {
	e := export.PHPExporter{}
	if e.ID() != "php" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "php")
	}
	if e.Name() != "PHP" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "PHP")
	}
}

func TestPHPExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.PHPExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "$BANNER = [\n  '##',\n  '  ',\n];\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestPHPExporterEscapesSingleQuote(t *testing.T) {
	lines := []string{`it's $5`}
	out, err := export.PHPExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `it\'s $5`) {
		t.Fatalf("expected escaped single quote with $ intact, got: %s", out)
	}
}
