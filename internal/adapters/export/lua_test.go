package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestLuaExporterIDAndName(t *testing.T) {
	e := export.LuaExporter{}
	if e.ID() != "lua" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "lua")
	}
	if e.Name() != "Lua" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "Lua")
	}
}

func TestLuaExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.LuaExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "local BANNER = {\n  \"##\",\n  \"  \",\n}\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestLuaExporterEscapesQuotesAndBackslashes(t *testing.T) {
	lines := []string{`say "hi"`, `back\slash`}
	out, err := export.LuaExporter{}.Export(lines)
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
