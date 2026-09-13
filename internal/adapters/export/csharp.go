package export

import "strings"

// CSharpExporter genera: string[] BANNER = {"linea1", "linea2"};.
type CSharpExporter struct{}

func (CSharpExporter) ID() string   { return "csharp" }
func (CSharpExporter) Name() string { return "C#" }

func (CSharpExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("string[] BANNER = {\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("};\n")
	return b.String(), nil
}
