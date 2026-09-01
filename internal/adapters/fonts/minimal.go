package fonts

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// minimalPatterns reutiliza los caracteres soportados por blockPatterns
// (para mantener el mismo charset y el mismo comportamiento de
// fallback) pero renderiza cada carácter como una única fila compacta,
// dando un estilo "minimal" bien diferenciado de block/big/banner.
var minimalPatterns = buildMinimalPatterns()

func buildMinimalPatterns() map[rune][]string {
	patterns := make(map[rune][]string, len(blockPatterns))
	for r := range blockPatterns {
		if r == ' ' {
			patterns[r] = []string{" "}
			continue
		}
		patterns[r] = []string{string(r)}
	}
	return patterns
}

// NewMinimalFont es una fuente de una sola fila, pensada para banners
// pequeños e inline.
func NewMinimalFont() domain.Font {
	return domain.Font{
		ID:         "minimal",
		Name:       "Minimal",
		Height:     1,
		Characters: buildCharacters(minimalPatterns),
		Fallback:   '?',
	}
}
