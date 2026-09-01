package export

import "strings"

// JavaScriptExporter genera: const BANNER = ['linea1', 'linea2', ...];
type JavaScriptExporter struct{}

func (JavaScriptExporter) ID() string   { return "javascript" }
func (JavaScriptExporter) Name() string { return "JavaScript" }

func (JavaScriptExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("const BANNER = [\n")
	for _, line := range lines {
		escaped := escapeBacktick(escapeSingleQuoted(line))
		b.WriteString("  '")
		b.WriteString(escaped)
		b.WriteString("',\n")
	}
	b.WriteString("];\n")
	return b.String(), nil
}
