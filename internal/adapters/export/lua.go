package export

import "strings"

// LuaExporter genera: local BANNER = {"linea1", "linea2"}.
type LuaExporter struct{}

func (LuaExporter) ID() string   { return "lua" }
func (LuaExporter) Name() string { return "Lua" }

func (LuaExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("local BANNER = {\n")
	for _, line := range lines {
		b.WriteString("  \"")
		b.WriteString(escapeDoubleQuoted(line))
		b.WriteString("\",\n")
	}
	b.WriteString("}\n")
	return b.String(), nil
}
