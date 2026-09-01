package fonts

import (
	"sort"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

// Registry es un adaptador secundario en memoria que implementa
// ports.FontRepository. El MVP incluye las 4 fuentes requeridas por el
// spec (Block, Big, Banner, Minimal); agregar una nueva fuente futura
// solo requiere sumarla acá, sin tocar el core.
type Registry struct {
	fonts map[string]domain.Font
}

var _ ports.FontRepository = (*Registry)(nil)

// NewRegistry construye el registro con las fuentes del MVP.
func NewRegistry() *Registry {
	all := []domain.Font{
		NewBlockFont(),
		NewBigFont(),
		NewBannerFont(),
		NewMinimalFont(),
	}
	byID := make(map[string]domain.Font, len(all))
	for _, f := range all {
		byID[f.ID] = f
	}
	return &Registry{fonts: byID}
}

func (r *Registry) Get(id string) (domain.Font, bool) {
	f, ok := r.fonts[id]
	return f, ok
}

func (r *Registry) List() []domain.FontInfo {
	infos := make([]domain.FontInfo, 0, len(r.fonts))
	for _, f := range r.fonts {
		infos = append(infos, domain.FontInfo{ID: f.ID, Name: f.Name})
	}
	sort.Slice(infos, func(i, j int) bool { return infos[i].ID < infos[j].ID })
	return infos
}
