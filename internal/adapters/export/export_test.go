package export_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestTxtExporterPreservesContent(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.TxtExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "##\n  " {
		t.Fatalf("unexpected txt output: %q", out)
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

func TestPythonExporterEscapesDoubleQuotes(t *testing.T) {
	lines := []string{`say "hi"`}
	out, err := export.PythonExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `say \"hi\"`) {
		t.Fatalf("expected escaped double quotes, got: %s", out)
	}
}

func TestJSONExporterProducesValidJSONWithUnicode(t *testing.T) {
	lines := []string{"█████", "line with \"quotes\" and \\backslash", ""}
	out, err := export.JSONExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var parsed []string
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("expected valid JSON, got error %v for: %s", err, out)
	}
	if len(parsed) != len(lines) {
		t.Fatalf("expected %d lines round-tripped, got %d", len(lines), len(parsed))
	}
	for i := range lines {
		if parsed[i] != lines[i] {
			t.Fatalf("line %d mismatch: got %q want %q", i, parsed[i], lines[i])
		}
	}
}

func TestRegistryListsAllRequiredFormats(t *testing.T) {
	reg := export.NewRegistry()
	want := []string{"txt", "javascript", "typescript", "rust", "python", "json"}
	for _, id := range want {
		if _, ok := reg.Get(id); !ok {
			t.Fatalf("expected exporter %q to be registered", id)
		}
	}
	if len(reg.List()) != len(want) {
		t.Fatalf("expected %d exporters, got %d", len(want), len(reg.List()))
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

func TestAllExportersReportIDAndName(t *testing.T) {
	reg := export.NewRegistry()
	want := map[string]string{
		"txt":        "Plain Text",
		"javascript": "JavaScript",
		"typescript": "TypeScript",
		"rust":       "Rust",
		"python":     "Python",
		"json":       "JSON",
	}
	for id, name := range want {
		e, ok := reg.Get(id)
		if !ok {
			t.Fatalf("exporter %q not registered", id)
		}
		if e.ID() != id {
			t.Fatalf("ID(): got %q, want %q", e.ID(), id)
		}
		if e.Name() != name {
			t.Fatalf("Name() for %q: got %q, want %q", id, e.Name(), name)
		}
	}
}

func TestRegistryGetUnknownFormatReturnsFalse(t *testing.T) {
	reg := export.NewRegistry()
	if _, ok := reg.Get("no-such-format"); ok {
		t.Fatal("expected ok=false for unknown format")
	}
}
