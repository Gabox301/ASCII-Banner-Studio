package export

import "strings"

// ShellExporter genera un array bash: BANNER=('linea1' 'linea2').
type ShellExporter struct{}

func (ShellExporter) ID() string   { return "shell" }
func (ShellExporter) Name() string { return "Shell" }

func (ShellExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("BANNER=(\n")
	for _, line := range lines {
		b.WriteString("  '")
		b.WriteString(escapeShellSingleQuoted(line))
		b.WriteString("'\n")
	}
	b.WriteString(")\n")
	return b.String(), nil
}
