package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestSwiftExporterIDAndName(t *testing.T) {
	e := export.SwiftExporter{}
	if e.ID() != "swift" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "swift")
	}
	if e.Name() != "Swift" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Swift")
	}
}

func TestSwiftExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.SwiftExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "let BANNER = [\n    \"##\",\n    \"  \",\n]\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestSwiftExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.SwiftExporter{}.Export(lines)
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
