package export

import "strings"

// PHPExporter genera: $BANNER = ['linea1', 'linea2']; con comillas simples
// para evitar interpolar $variables del arte.
type PHPExporter struct{}

func (PHPExporter) ID() string   { return "php" }
func (PHPExporter) Name() string { return "PHP" }

func (PHPExporter) Export(lines []string) (string, error) {
	var b strings.Builder
	b.WriteString("$BANNER = [\n")
	for _, line := range lines {
		b.WriteString("  '")
		b.WriteString(escapeSingleQuoted(line))
		b.WriteString("',\n")
	}
	b.WriteString("];\n")
	return b.String(), nil
}
