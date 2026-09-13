package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestJavaExporterIDAndName(t *testing.T) {
	e := export.JavaExporter{}
	if e.ID() != "java" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "java")
	}
	if e.Name() != "Java" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Java")
	}
}

func TestJavaExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.JavaExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "String[] BANNER = {\n    \"##\",\n    \"  \",\n};\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestJavaExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.JavaExporter{}.Export(lines)
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
