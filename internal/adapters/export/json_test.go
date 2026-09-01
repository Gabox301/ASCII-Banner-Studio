package export_test

import (
	"encoding/json"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestJSONExporterID(t *testing.T) {
	e := export.JSONExporter{}
	if e.ID() != "json" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "json")
	}
}

func TestJSONExporterName(t *testing.T) {
	e := export.JSONExporter{}
	if e.Name() != "JSON" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "JSON")
	}
}

func TestJSONExporterProducesValidJSON(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.JSONExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var parsed []string
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("expected valid JSON, got error %v for: %s", err, out)
	}
	if len(parsed) != len(lines) {
		t.Fatalf("expected %d lines, got %d", len(lines), len(parsed))
	}
	for i := range lines {
		if parsed[i] != lines[i] {
			t.Fatalf("line %d mismatch: got %q, want %q", i, parsed[i], lines[i])
		}
	}
}

func TestJSONExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.JSONExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var parsed []string
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("expected valid JSON, got error %v for: %s", err, out)
	}
	if parsed[0] != `say "hi"` {
		t.Fatalf("expected original quote, got %q", parsed[0])
	}
	if parsed[1] != `back\slash` {
		t.Fatalf("expected original backslash, got %q", parsed[1])
	}
}

func TestJSONExporterHandlesEmptyLines(t *testing.T) {
	lines := []string{"", "text", ""}
	out, err := export.JSONExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var parsed []string
	if err := json.Unmarshal([]byte(out), &parsed); err != nil {
		t.Fatalf("expected valid JSON, got error %v for: %s", err, out)
	}
	if len(parsed) != len(lines) {
		t.Fatalf("expected %d lines, got %d", len(lines), len(parsed))
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
