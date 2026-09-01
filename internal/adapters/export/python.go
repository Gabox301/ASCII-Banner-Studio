package export

import "strings"

// PythonExporter genera: BANNER = ["linea1", "linea2"]
type PythonExporter struct{}

func (PythonExporter) ID() string   { return "python" }
func (PythonExporter) Name() string { return "Python" }

func (PythonExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("BANNER = [\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("]\n")
	return b.String(), nil
}
