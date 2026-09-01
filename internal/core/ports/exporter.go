package ports

// Exporter es un puerto secundario para transformar el banner generado
// (como líneas de texto) en distintos formatos de salida (txt, js, ts,
// rust, python, json, ...). Cada implementación vive aislada del
// renderer, tal como exige el spec original.
type Exporter interface {
	ID() string
	Name() string
	Export(lines []string) (string, error)
}

// ExporterRepository agrupa exporters disponibles por id.
type ExporterRepository interface {
	Get(id string) (Exporter, bool)
	List() []Exporter
}
