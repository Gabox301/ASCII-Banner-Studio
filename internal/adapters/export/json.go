package export

import "encoding/json"

// JSONExporter genera un array JSON de líneas, delegando el escapado
// (comillas, backslashes, Unicode) en encoding/json para garantizar
// sintaxis siempre válida.
type JSONExporter struct{}

func (JSONExporter) ID() string   { return "json" }
func (JSONExporter) Name() string { return "JSON" }

func (JSONExporter) Export(lines []string) (string, error) {
	data, err := json.MarshalIndent(lines, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}
