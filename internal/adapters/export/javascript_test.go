package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestJavaScriptExporterID(t *testing.T) {
	e := export.JavaScriptExporter{}
	if e.ID() != "javascript" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "javascript")
	}
}

func TestJavaScriptExporterName(t *testing.T) {
	e := export.JavaScriptExporter{}
	if e.Name() != "JavaScript" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "JavaScript")
	}
}

func TestJavaScriptExporterMirrorBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.JavaScriptExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "const BANNER = [\n  '##',\n  '  ',\n];\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestJavaScriptExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`it's a "test"`, `back\slash`}
	out, err := export.JavaScriptExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `it\'s a "test"`) {
		t.Fatalf("expected escaped single quote, got: %s", out)
	}
	if !strings.Contains(out, `back\\slash`) {
		t.Fatalf("expected escaped backslash, got: %s", out)
	}
}

func TestJavaScriptExporterEscapesBackticks(t *testing.T) {
	lines := []string{"`template`"}
	out, err := export.JavaScriptExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "\\`template\\`") {
		t.Fatalf("expected escaped backticks, got: %s", out)
	}
}
