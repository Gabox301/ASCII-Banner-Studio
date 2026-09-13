package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestGoExporterIDAndName(t *testing.T) {
	e := export.GoExporter{}
	if e.ID() != "go" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "go")
	}
	if e.Name() != "Go" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Go")
	}
}

func TestGoExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.GoExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "var BANNER = []string{\n    \"##\",\n    \"  \",\n}\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestGoExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.GoExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `say \"hi\"`) {
		t.Fatalf("expected escaped double quote, got: %s", out)
	}
	if !strings.Contains(out, `back\\slash`) {
		t.Fatalf("expected escaped backslash, got: %s", out)
	}
}
