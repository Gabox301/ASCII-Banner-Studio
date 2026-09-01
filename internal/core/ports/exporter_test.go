package ports_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

type mockExporter struct{}

func (mockExporter) ID() string   { return "mock" }
func (mockExporter) Name() string { return "Mock" }
func (mockExporter) Export(lines []string) (string, error) {
	return "", nil
}

type mockExporterRepo struct{}

func (mockExporterRepo) Get(id string) (ports.Exporter, bool) {
	return nil, false
}
func (mockExporterRepo) List() []ports.Exporter {
	return nil
}

func TestExporterInterface(t *testing.T) {
	var _ ports.Exporter = mockExporter{}
	var _ ports.ExporterRepository = mockExporterRepo{}
	t.Log("Exporter and ExporterRepository interfaces are satisfied by mock implementations")
}
