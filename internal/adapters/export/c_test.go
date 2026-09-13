package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestCExporterIDAndName(t *testing.T) {
	e := export.CExporter{}
	if e.ID() != "c" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "c")
	}
	if e.Name() != "C/C++" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "C/C++")
	}
}

func TestCExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.CExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "const char *BANNER[] = {\n    \"##\",\n    \"  \",\n};\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestCExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.CExporter{}.Export(lines)
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
