// Package service contiene los casos de uso de la aplicación. Depende
// únicamente de domain y ports: nunca de un adaptador concreto. Esto es
// lo que permite reutilizar el motor tanto desde la CLI actual como
// desde cualquier interfaz futura (equivalente al ascii-core del spec
// original, que no debía depender de Tauri).
package service

import (
	"strings"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

// BannerService implementa ports.BannerGenerator.
type BannerService struct {
	fonts ports.FontRepository
}

// NewBannerService crea el servicio inyectando el puerto de fuentes.
func NewBannerService(fonts ports.FontRepository) *BannerService {
	return &BannerService{fonts: fonts}
}

var _ ports.BannerGenerator = (*BannerService)(nil)

// ListFonts delega en el repositorio de fuentes.
func (s *BannerService) ListFonts() []domain.FontInfo {
	return s.fonts.List()
}

// Generate implementa el algoritmo descripto en el spec:
//  1. valida longitud;
//  2. normaliza (trim / uppercase);
//  3. resuelve el glyph de cada carácter (con fallback, sin panics);
//  4. concatena horizontalmente aplicando spacing;
//  5. soporta múltiples líneas de entrada (separadas por \n), cada una
//     generando su propio bloque, alineados entre sí;
//  6. devuelve el resultado como string.
func (s *BannerService) Generate(text string, opts domain.BannerOptions) (string, error) {
	if len([]rune(text)) > domain.MaxTextLength {
		return "", domain.ErrTextTooLong
	}

	font, ok := s.fonts.Get(opts.Font)
	if !ok {
		return "", domain.ErrFontNotFound
	}

	if opts.Trim {
		text = strings.TrimSpace(text)
	}
	if opts.Uppercase {
		text = strings.ToUpper(text)
	}
	if text == "" {
		return "", nil
	}

	spacing := opts.Spacing
	if spacing < 0 {
		spacing = 0
	}

	inputLines := strings.Split(text, "\n")
	blocks := make([][]string, 0, len(inputLines))
	maxWidth := 0

	for _, inputLine := range inputLines {
		block := renderLine(font, inputLine, spacing)
		for _, row := range block {
			if w := len([]rune(row)); w > maxWidth {
				maxWidth = w
			}
		}
		blocks = append(blocks, block)
	}

	align := opts.Align
	if align == "" {
		align = domain.AlignLeft
	}

	var out []string
	for i, block := range blocks {
		if i > 0 {
			out = append(out, alignLine("", maxWidth, align)) // separador entre bloques multilinea
		}
		for _, row := range block {
			out = append(out, alignLine(row, maxWidth, align))
		}
	}

	return strings.Join(out, "\n"), nil
}

// renderLine aplica el pseudocódigo del spec para una única línea de
// texto: por cada fila del glyph, concatena el glyph de cada carácter
// más el spacing.
func renderLine(font domain.Font, line string, spacing int) []string {
	rows := make([]strings.Builder, font.Height)
	runes := []rune(line)

	if len(runes) == 0 {
		blank := make([]string, font.Height)
		for i := range blank {
			blank[i] = ""
		}
		return blank
	}

	pad := strings.Repeat(" ", spacing)

	for _, r := range runes {
		glyph := font.Glyph(r)
		for i := 0; i < font.Height; i++ {
			if i < len(glyph) {
				rows[i].WriteString(glyph[i])
			}
			rows[i].WriteString(pad)
		}
	}

	result := make([]string, font.Height)
	for i := range rows {
		result[i] = rows[i].String()
	}
	return result
}

// alignLine rellena una línea con espacios hasta maxWidth según la
// alineación pedida.
func alignLine(line string, maxWidth int, align domain.Alignment) string {
	width := len([]rune(line))
	if width >= maxWidth {
		return line
	}
	gap := maxWidth - width

	switch align {
	case domain.AlignRight:
		return strings.Repeat(" ", gap) + line
	case domain.AlignCenter:
		left := gap / 2
		right := gap - left
		return strings.Repeat(" ", left) + line + strings.Repeat(" ", right)
	default: // AlignLeft
		return line + strings.Repeat(" ", gap)
	}
}
