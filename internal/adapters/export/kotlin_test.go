package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestKotlinExporterIDAndName(t *testing.T) {
	e := export.KotlinExporter{}
	if e.ID() != "kotlin" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "kotlin")
	}
	if e.Name() != "Kotlin" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Kotlin")
	}
}

func TestKotlinExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.KotlinExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "val BANNER = listOf(\n    \"##\",\n    \"  \",\n)\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestKotlinExporterEscapesDollarTemplate(t *testing.T) {
	lines := []string{`price $5 and ${x}`}
	out, err := export.KotlinExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `price ${'$'}5 and ${'$'}{x}`) {
		t.Fatalf("expected escaped dollar signs, got: %s", out)
	}
}
