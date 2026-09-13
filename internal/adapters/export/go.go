package export

import "strings"

// GoExporter genera: var BANNER = []string{"linea1", "linea2"}.
type GoExporter struct{}

func (GoExporter) ID() string   { return "go" }
func (GoExporter) Name() string { return "Go" }

func (GoExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("var BANNER = []string{\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("}\n")
	return b.String(), nil
}
