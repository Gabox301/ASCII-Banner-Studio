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

func TestSwiftExporterEscapesInterpolation(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{name: "interpolation opener is neutralized", lines: []string{`\(name)`}, want: `    "\\(name)",` + "\n"},
		{name: "ascii art backslash paren is neutralized", lines: []string{`\(`}, want: `    "\\(",` + "\n"},
		{name: "plain parens are untouched", lines: []string{`(x)`}, want: `    "(x)",` + "\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := export.SwiftExporter{}.Export(tt.lines)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(out, tt.want) {
				t.Fatalf("expected output to contain %q, got: %s", tt.want, out)
			}
		})
	}
}
