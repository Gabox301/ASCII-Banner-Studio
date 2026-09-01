package fonts

import (
	"strings"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

// NewBigFont deriva sus glyphs de blockPatterns escalando cada pixel
// x2 en ambos ejes. Vive en su propio archivo y no modifica el
// algoritmo del core: solo aporta un FontRepository entry más.
func NewBigFont() domain.Font {
	const scale = 2
	scaled := make(map[rune][]string, len(blockPatterns))
	for r, lines := range blockPatterns {
		scaled[r] = scaleGlyph(lines, scale)
	}
	return domain.Font{
		ID:         "big",
		Name:       "Big",
		Height:     5 * scale,
		Characters: buildCharacters(scaled),
		Fallback:   '?',
	}
}

// scaleGlyph expande un glyph reemplazando cada pixel por un bloque de
// scale x scale, preservando el ancho consistente por fila que exige
// domain.Glyph.
func scaleGlyph(lines []string, scale int) []string {
	out := make([]string, 0, len(lines)*scale)
	for _, line := range lines {
		var wide strings.Builder
		for _, c := range line {
			wide.WriteString(strings.Repeat(string(c), scale))
		}
		row := wide.String()
		for i := 0; i < scale; i++ {
			out = append(out, row)
		}
	}
	return out
}
