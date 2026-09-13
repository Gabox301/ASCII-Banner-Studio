package export

import "strings"

// CExporter genera: const char *BANNER[] = {"linea1", "linea2"}; (vale
// para C y C++).
type CExporter struct{}

func (CExporter) ID() string   { return "c" }
func (CExporter) Name() string { return "C/C++" }

func (CExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("const char *BANNER[] = {\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("};\n")
	return b.String(), nil
}
