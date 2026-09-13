package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestRubyExporterIDAndName(t *testing.T) {
	e := export.RubyExporter{}
	if e.ID() != "ruby" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "ruby")
	}
	if e.Name() != "Ruby" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Ruby")
	}
}

func TestRubyExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.RubyExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "BANNER = [\n  \"##\",\n  \"  \",\n]\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestRubyExporterEscapesInterpolation(t *testing.T) {
	lines := []string{`#{name}`}
	out, err := export.RubyExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `\#{name}`) {
		t.Fatalf("expected escaped interpolation, got: %s", out)
	}
}
