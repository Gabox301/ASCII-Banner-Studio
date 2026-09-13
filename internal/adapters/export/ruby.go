package export

import "strings"

// RubyExporter genera: BANNER = ["linea1", "linea2"].
type RubyExporter struct{}

func (RubyExporter) ID() string   { return "ruby" }
func (RubyExporter) Name() string { return "Ruby" }

func (RubyExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("BANNER = [\n")
	for _, line := range lines {
		b.WriteString("  \"")
		b.WriteString(escapeRubyDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("]\n")
	return b.String(), nil
}
