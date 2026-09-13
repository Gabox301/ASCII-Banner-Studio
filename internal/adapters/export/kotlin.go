package export

import "strings"

// KotlinExporter genera: val BANNER = listOf("linea1", "linea2").
type KotlinExporter struct{}

func (KotlinExporter) ID() string   { return "kotlin" }
func (KotlinExporter) Name() string { return "Kotlin" }

func (KotlinExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("val BANNER = listOf(\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeKotlin(line))
		b.WriteString("\",\n")
	}
	b.WriteString(")\n")
	return b.String(), nil
}
