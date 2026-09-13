package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestCSharpExporterIDAndName(t *testing.T) {
	e := export.CSharpExporter{}
	if e.ID() != "csharp" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "csharp")
	}
	if e.Name() != "C#" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "C#")
	}
}

func TestCSharpExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.CSharpExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "string[] BANNER = {\n    \"##\",\n    \"  \",\n};\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestCSharpExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.CSharpExporter{}.Export(lines)
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
