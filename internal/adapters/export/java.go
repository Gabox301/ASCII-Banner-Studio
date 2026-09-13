package export

import "strings"

// JavaExporter genera: String[] BANNER = {"linea1", "linea2"};.
type JavaExporter struct{}

func (JavaExporter) ID() string   { return "java" }
func (JavaExporter) Name() string { return "Java" }

func (JavaExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("String[] BANNER = {\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("};\n")
	return b.String(), nil
}
