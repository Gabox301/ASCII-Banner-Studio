package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestTypeScriptExporterID(t *testing.T) {
	e := export.TypeScriptExporter{}
	if e.ID() != "typescript" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "typescript")
	}
}

func TestTypeScriptExporterName(t *testing.T) {
	e := export.TypeScriptExporter{}
	if e.Name() != "TypeScript" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "TypeScript")
	}
}

func TestTypeScriptExporterMirrorBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.TypeScriptExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "export const BANNER: string[] = [\n  '##',\n  '  ',\n];\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestTypeScriptExporterProducesArray(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.TypeScriptExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "export const BANNER: string[] = [") {
		t.Fatalf("expected typescript wrapper, got: %s", out)
	}
	if !strings.Contains(out, "'##',") {
		t.Fatalf("expected single-quoted row, got: %s", out)
	}
	if !strings.HasSuffix(out, "];\n") {
		t.Fatalf("expected closing `];`, got: %s", out)
	}
}
