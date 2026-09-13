package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
)

func TestPowerShellExporterIDAndName(t *testing.T) {
	e := export.PowerShellExporter{}
	if e.ID() != "powershell" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "powershell")
	}
	if e.Name() != "PowerShell" {
		t.Fatalf("Name: got %q, want %q", e.Name(), "PowerShell")
	}
}

func TestPowerShellExporterBasicOutput(t *testing.T) {
	lines := []string{"##", "  "}
	out, err := export.PowerShellExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "$BANNER = @(\n  '##',\n  '  '\n)\n"
	if out != expected {
		t.Fatalf("unexpected output: got %q, want %q", out, expected)
	}
}

func TestPowerShellExporterEscapesSingleQuote(t *testing.T) {
	lines := []string{`it's`}
	out, err := export.PowerShellExporter{}.Export(lines)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `it''s`) {
		t.Fatalf("expected doubled single quote, got: %s", out)
	}
}
