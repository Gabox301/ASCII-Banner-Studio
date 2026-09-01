package ports

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// BannerGenerator es el puerto primario (driving port): la forma en que
// cualquier adaptador de entrada (CLI, futura GUI, tests) invoca el
// caso de uso principal del sistema.
type BannerGenerator interface {
	Generate(text string, opts domain.BannerOptions) (string, error)
	ListFonts() []domain.FontInfo
}
