package fonts

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// NewBlockFont es la fuente por defecto: bloques sólidos de 5 filas.
func NewBlockFont() domain.Font {
	return domain.Font{
		ID:         "block",
		Name:       "Block",
		Height:     5,
		Characters: buildCharacters(blockPatterns),
		Fallback:   '?',
	}
}
