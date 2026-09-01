package export_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

func TestRegistryGet(t *testing.T) {
	reg := export.NewRegistry()
	e, ok := reg.Get("txt")
	if !ok {
		t.Fatal("expected ok=true for known format")
	}
	if e.ID() != "txt" {
		t.Fatalf("ID: got %q, want %q", e.ID(), "txt")
	}
}

func TestRegistryGetUnknownFormat(t *testing.T) {
	reg := export.NewRegistry()
	if _, ok := reg.Get("no-such-format"); ok {
		t.Fatal("expected ok=false for unknown format")
	}
}

func TestRegistryList(t *testing.T) {
	reg := export.NewRegistry()
	list := reg.List()
	if len(list) != 6 {
		t.Fatalf("expected 6 exporters, got %d", len(list))
	}
}

func TestRegistryListOrder(t *testing.T) {
	reg := export.NewRegistry()
	ids := make([]string, len(reg.List()))
	for i, e := range reg.List() {
		ids[i] = e.ID()
	}
	want := []string{"txt", "javascript", "typescript", "rust", "python", "json"}
	for i := range want {
		if ids[i] != want[i] {
			t.Fatalf("index %d: got %q, want %q", i, ids[i], want[i])
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

func TestRegistryAllExportersReportIDAndName(t *testing.T) {
	reg := export.NewRegistry()
	want := map[string]string{
		"txt":        "Plain Text",
		"javascript": "JavaScript",
		"typescript": "TypeScript",
		"rust":       "Rust",
		"python":     "Python",
		"json":       "JSON",
	}
	for _, e := range reg.List() {
		if name, ok := want[e.ID()]; !ok || name != e.Name() {
			t.Fatalf("exporter %q: got Name()=%q, want %q", e.ID(), e.Name(), want[e.ID()])
		}
	}
}

func TestNewRegistryImplementsExporterRepository(t *testing.T) {
	var _ ports.ExporterRepository = export.NewRegistry()
	_ = strings.TrimSpace("")
	t.Log("Registry implements ExporterRepository interface")
}
