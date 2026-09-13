package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestShellExporterIDAndName(t *testing.T) {
	e := export.ShellExporter{}
	if e.ID() != "shell" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "shell")
	}
	if e.Name() != "Shell" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Shell")
	}
}

func TestShellExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.ShellExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "BANNER=(\n  '##'\n  '  '\n)\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestShellExporterEscapesSingleQuote(t *testing.T) {
	lines := []string{`it's`}
	out, err := export.ShellExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `it'\''s`) {
		t.Fatalf("expected shell-escaped single quote, got: %s", out)
	}
}
