package export

import "github.com/gabriel/ascii-banner-studio/internal/core/ports"

// Registry agrupa todos los exporters disponibles del MVP.
type Registry struct {
	exporters map[string]ports.Exporter
	order     []string
}

var _ ports.ExporterRepository = (*Registry)(nil)

// NewRegistry construye el registro con los 6 formatos requeridos por
// el spec: TXT, JavaScript, TypeScript, Rust, Python y JSON.
func NewRegistry() *Registry {
	list := []ports.Exporter{
		TxtExporter{},
		JavaScriptExporter{},
		TypeScriptExporter{},
		RustExporter{},
		PythonExporter{},
		JSONExporter{},
	}
	byID := make(map[string]ports.Exporter, len(list))
	order := make([]string, 0, len(list))
	for _, e := range list {
		byID[e.ID()] = e
		order = append(order, e.ID())
	}
	return &Registry{exporters: byID, order: order}
}

func (r *Registry) Get(id string) (ports.Exporter, bool) {
	e, ok := r.exporters[id]
	return e, ok
}

func (r *Registry) List() []ports.Exporter {
	out := make([]ports.Exporter, 0, len(r.order))
	for _, id := range r.order {
		out = append(out, r.exporters[id])
	}
	return out
}
