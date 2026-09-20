package export

import "strings"

// SwiftExporter genera: let BANNER = ["linea1", "linea2"].
type SwiftExporter struct{}

func (SwiftExporter) ID() string   { return "swift" }
func (SwiftExporter) Name() string { return "Swift" }

func (SwiftExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("let BANNER = [\n")
	for _, line := range lines {
		b.WriteString("    \"")
		b.WriteString(escapeSwift(line))
		b.WriteString("\",\n")
	}
	b.WriteString("]\n")
	return b.String(), nil
}
