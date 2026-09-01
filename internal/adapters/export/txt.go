package export

import "strings"

// TxtExporter devuelve el ASCII puro, sin modificar espacios ni saltos
// de línea (equivalente al botón Copy / export TXT del spec).
type TxtExporter struct{}

func (TxtExporter) ID() string   { return "txt" }
func (TxtExporter) Name() string { return "Plain Text" }

func (TxtExporter) Export(lines []string) (string, error) {
	return strings.Join(lines, "\n"), nil
}
