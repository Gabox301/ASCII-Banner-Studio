package fonts

import (
	"strings"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

// bannerFill es el carácter Unicode usado para el "peso visual" de la
// fuente banner, similar a los banners generados con bloques llenos
// (ej. el ejemplo "SKILLINDEX" del spec original).
const bannerFill = "█"

// NewBannerFont deriva sus glyphs de blockPatterns sustituyendo '#' por
// un bloque Unicode sólido y agregando un espacio de margen a cada
// lado del glyph.
func NewBannerFont() domain.Font {
	styled := make(map[rune][]string, len(blockPatterns))
	for r, lines := range blockPatterns {
		styled[r] = styleGlyph(lines)
	}
	return domain.Font{
		ID:         "banner",
		Name:       "Banner",
		Height:     5,
		Characters: buildCharacters(styled),
		Fallback:   '?',
	}
}

func styleGlyph(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		replaced := strings.ReplaceAll(line, "#", bannerFill)
		out[i] = " " + replaced + " "
	}
	return out
}
