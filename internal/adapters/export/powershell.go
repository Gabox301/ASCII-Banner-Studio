package export

import "strings"

// PowerShellExporter genera: $BANNER = @('linea1', 'linea2').
type PowerShellExporter struct{}

func (PowerShellExporter) ID() string   { return "powershell" }
func (PowerShellExporter) Name() string { return "PowerShell" }

func (PowerShellExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("$BANNER = @(\n")
	for i, line := range lines {
		b.WriteString("  '")
		b.WriteString(escapePowerShellSingleQuoted(line))
		b.WriteString("'")
		if i < len(lines)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString(")\n")
	return b.String(), nil
}
