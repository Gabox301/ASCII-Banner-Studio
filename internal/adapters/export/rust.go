package export

import "strings"

// RustExporter genera: const BANNER: &[&str] = &["linea1", "linea2"];
type RustExporter struct{}

func (RustExporter) ID() string   { return "rust" }
func (RustExporter) Name() string { return "Rust" }

func (RustExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("const BANNER: &[&str] = &[\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("];\n")
	return b.String(), nil
}
