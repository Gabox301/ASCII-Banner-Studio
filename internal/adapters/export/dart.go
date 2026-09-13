package export

import "strings"

// DartExporter genera: const BANNER = ['linea1', 'linea2'];.
type DartExporter struct{}

func (DartExporter) ID() string   { return "dart" }
func (DartExporter) Name() string { return "Dart" }

func (DartExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("const BANNER = [\n")
	for _, line := range lines {
		b.WriteString("  '")
		b.WriteString(escapeDartSingleQuoted(line))
		b.WriteString("',\n")
	}
	b.WriteString("];\n")
	return b.String(), nil
}
