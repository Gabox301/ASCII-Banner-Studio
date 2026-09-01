package export

import "strings"

// TypeScriptExporter genera: export const BANNER: string[] = [...]
type TypeScriptExporter struct{}

func (TypeScriptExporter) ID() string   { return "typescript" }
func (TypeScriptExporter) Name() string { return "TypeScript" }

func (TypeScriptExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("export const BANNER: string[] = [\n")
	for _, line := range lines {
		escaped := escapeBacktick(escapeSingleQuoted(line))
		b.WriteString("  '")
		b.WriteString(escaped)
		b.WriteString("',\n")
	}
	b.WriteString("];\n")
	return b.String(), nil
}
