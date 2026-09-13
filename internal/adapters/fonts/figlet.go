package fonts

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// buildFigletCharacters convierte patrones FIGlet (mapa rune -> filas) en el
// mapa que espera domain.Font, preservando minúsculas distintas cuando la
// fuente las define y mapeando a mayúsculas solo si faltan (a diferencia de
// buildCharacters, que siempre sobrescribe y es no-determinista con mapas que
// ya traen ambas cajas).
func buildFigletCharacters(patterns map[rune][]string) map[rune]domain.Glyph {
	chars := make(map[rune]domain.Glyph, len(patterns)+26)
	for r, lines := range patterns {
		glyph := make(domain.Glyph, len(lines))
		copy(glyph, lines)
		chars[r] = glyph
	}
	for r := 'A'; r <= 'Z'; r++ {
		lower := r + ('a' - 'A')
		if _, ok := chars[lower]; !ok {
			if g, ok := chars[r]; ok {
				chars[lower] = g
			}
		}
	}
	return chars
}
