// Package domain contiene las entidades y reglas de negocio puras del
// motor de banners ASCII. No importa nada de infraestructura (ni Tauri,
// ni CLI, ni frameworks): es el núcleo del hexágono.
package domain

// Alignment representa la alineación horizontal del banner generado.
type Alignment string

const (
	AlignLeft   Alignment = "left"
	AlignCenter Alignment = "center"
	AlignRight  Alignment = "right"
)

// BannerOptions agrupa las opciones de generación, equivalentes a
// BannerOptions del spec original.
type BannerOptions struct {
	Font      string
	Spacing   int
	Align     Alignment
	Uppercase bool
	Trim      bool
}

// Glyph es la representación de un carácter como una matriz de líneas.
// Todas las líneas de un mismo Glyph deben tener el mismo ancho.
type Glyph []string

// Font representa una fuente ASCII: nombre, altura fija y el mapa de
// caracteres soportados. Fallback indica el rune a usar cuando el
// carácter de entrada no está soportado (nunca debe hacer panic).
type Font struct {
	ID         string
	Name       string
	Height     int
	Characters map[rune]Glyph
	Fallback   rune
}

// FontInfo es la proyección liviana usada para listar fuentes
// disponibles (equivalente a list_fonts() del spec).
type FontInfo struct {
	ID   string
	Name string
}

// Glyph devuelve la matriz de líneas para un carácter, resolviendo el
// fallback configurado si el carácter no está soportado, y garantizando
// que jamás se dispare un panic ni se devuelva un glyph de ancho
// inconsistente.
func (f Font) Glyph(r rune) Glyph {
	if g, ok := f.Characters[r]; ok {
		return g
	}
	if g, ok := f.Characters[f.Fallback]; ok {
		return g
	}
	// Último recurso: glyph en blanco con la altura correcta.
	blank := make(Glyph, f.Height)
	for i := range blank {
		blank[i] = ""
	}
	return blank
}

// Width devuelve el ancho (en runas) de un glyph, asumiendo que todas
// sus líneas tienen el mismo ancho (invariante que deben cumplir todas
// las fuentes del sistema).
func (g Glyph) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len([]rune(g[0]))
}
