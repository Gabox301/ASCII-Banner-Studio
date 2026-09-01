package ports

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// FontRepository es un puerto secundario (driven port): el núcleo lo
// necesita para resolver fuentes, pero no le importa de dónde vengan
// (memoria, JSON, editor de fuentes futuro, etc).
type FontRepository interface {
	Get(id string) (domain.Font, bool)
	List() []domain.FontInfo
}
