package fonts

import (
	"sort"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

// Registry es un adaptador secundario en memoria que implementa
// ports.FontRepository. Incluye las 4 fuentes base (Block, Big, Banner,
// Minimal) más 20 fuentes estilo FIGlet; agregar una nueva fuente futura
// solo requiere sumarla acá, sin tocar el core.
//
// Para evitar duplicaciones, las FIGlet "big" y "block" se registran como
// "big-figlet" y "block-figlet" (las IDs "big" y "block" ya existen con otro
// diseño). "banner3d" convive con "banner" porque son diseños distintos.
type Registry struct {
	fonts map[string]domain.Font
}

var _ ports.FontRepository = (*Registry)(nil)

// NewRegistry construye el registro con todas las fuentes disponibles.
func NewRegistry() *Registry {
	all := []domain.Font{
		NewBlockFont(),
		NewBigFont(),
		NewBannerFont(),
		NewMinimalFont(),
		NewStandardFont(),
		NewBigFigletFont(),
		NewSlantFont(),
		NewBlockFigletFont(),
		NewShadowFont(),
		NewBubbleFont(),
		NewDigitalFont(),
		NewStarwarsFont(),
		NewDoomFont(),
		NewScriptFont(),
		NewBanner3DFont(),
		NewIsometricFont(),
		NewLarry3DFont(),
		NewOgreFont(),
		NewGraffitiFont(),
		NewAnsiShadowFont(),
		NewColossalFont(),
		NewPyramidFont(),
		NewTinkerToyFont(),
		NewEpicFont(),
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
