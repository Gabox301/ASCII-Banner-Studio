package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestRustExporterID(t *testing.T) {
	e := export.RustExporter{}
	if e.ID() != "rust" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "rust")
	}
}

func TestRustExporterName(t *testing.T) {
	e := export.RustExporter{}
	if e.Name() != "Rust" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Rust")
	}
}

func TestRustExporterMirrorBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.RustExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "const BANNER: &[&str] = &[\n    \"##\",\n    \"  \",\n];\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestRustExporterEscapesDoubleQuotes(t *testing.T) {
	lines := []string{`say "hi"`}
	out, err := export.RustExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `say \"hi\"`) {
		t.Fatalf("expected escaped double quotes, got: %s", out)
	}
	if !strings.HasPrefix(out, "const BANNER: &[&str] = &[") {
		t.Fatalf("expected valid rust array header, got: %s", out)
	}
}
