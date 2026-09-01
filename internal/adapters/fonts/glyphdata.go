// Package fonts contiene los adaptadores secundarios que implementan
// ports.FontRepository. Cada fuente vive en su propio archivo y puede
// agregarse sin tocar el algoritmo del core (tal como pide el spec:
// "Cada fuente debe estar aislada y poder agregarse sin modificar el
// algoritmo principal").
package fonts

import "github.com/gabriel/ascii-banner-studio/internal/core/domain"

// blockPatterns define la matriz 5x5 base (carácter '#' = pixel
// encendido) para A-Z, 0-9, espacio y un subset de puntuación. Es la
// fuente de la que derivan "block", "big" (escalada x2) y "banner"
// (con relleno unicode).
//
// Los caracteres no incluidos aquí caen en el fallback '?' del motor:
// esto es intencional y cumple el requisito de que un carácter no
// soportado nunca rompa la aplicación.
var blockPatterns = map[rune][]string{
	'A': {" ### ", "#   #", "#####", "#   #", "#   #"},
	'B': {"#### ", "#   #", "#### ", "#   #", "#### "},
	'C': {" ####", "#    ", "#    ", "#    ", " ####"},
	'D': {"#### ", "#   #", "#   #", "#   #", "#### "},
	'E': {"#####", "#    ", "#### ", "#    ", "#####"},
	'F': {"#####", "#    ", "#### ", "#    ", "#    "},
	'G': {" ####", "#    ", "#  ##", "#   #", " ####"},
	'H': {"#   #", "#   #", "#####", "#   #", "#   #"},
	'I': {"#####", "  #  ", "  #  ", "  #  ", "#####"},
	'J': {"    #", "    #", "    #", "#   #", " ### "},
	'K': {"#   #", "#  # ", "###  ", "#  # ", "#   #"},
	'L': {"#    ", "#    ", "#    ", "#    ", "#####"},
	'M': {"#   #", "## ##", "# # #", "#   #", "#   #"},
	'N': {"#   #", "##  #", "# # #", "#  ##", "#   #"},
	'O': {" ### ", "#   #", "#   #", "#   #", " ### "},
	'P': {"#### ", "#   #", "#### ", "#    ", "#    "},
	'Q': {" ### ", "#   #", "#   #", "#  # ", " ## #"},
	'R': {"#### ", "#   #", "#### ", "#  # ", "#   #"},
	'S': {" ####", "#    ", " ### ", "    #", "#### "},
	'T': {"#####", "  #  ", "  #  ", "  #  ", "  #  "},
	'U': {"#   #", "#   #", "#   #", "#   #", " ### "},
	'V': {"#   #", "#   #", "#   #", " # # ", "  #  "},
	'W': {"#   #", "#   #", "# # #", "## ##", "#   #"},
	'X': {"#   #", " # # ", "  #  ", " # # ", "#   #"},
	'Y': {"#   #", " # # ", "  #  ", "  #  ", "  #  "},
	'Z': {"#####", "   # ", "  #  ", " #   ", "#####"},

	'0': {" ### ", "#   #", "#  ##", "##  #", " ### "},
	'1': {"  #  ", " ##  ", "  #  ", "  #  ", "#####"},
	'2': {" ### ", "#   #", "   # ", "  #  ", "#####"},
	'3': {"#####", "   # ", "  ## ", "    #", "#### "},
	'4': {"   ##", "  # #", " #  #", "#####", "    #"},
	'5': {"#####", "#    ", "#### ", "    #", "#### "},
	'6': {" ####", "#    ", "#### ", "#   #", " ### "},
	'7': {"#####", "    #", "   # ", "  #  ", "  #  "},
	'8': {" ### ", "#   #", " ### ", "#   #", " ### "},
	'9': {" ### ", "#   #", " ####", "    #", " ### "},

	' ': {"   ", "   ", "   ", "   ", "   "},
	'.': {"   ", "   ", "   ", "   ", " # "},
	',': {"   ", "   ", "   ", " # ", "#  "},
	'!': {" # ", " # ", " # ", "   ", " # "},
	'?': {" ### ", "#   #", "   # ", "  #  ", "  #  "},
	'-': {"     ", "     ", "#####", "     ", "     "},
	'_': {"     ", "     ", "     ", "     ", "#####"},
}

// buildCharacters convierte blockPatterns en el mapa que espera
// domain.Font, mapeando también minúsculas al mismo glyph que su
// mayúscula (una fuente en bloque no distingue caja).
func buildCharacters(patterns map[rune][]string) map[rune]domain.Glyph {
	chars := make(map[rune]domain.Glyph, len(patterns)*2)
	for r, lines := range patterns {
		glyph := make(domain.Glyph, len(lines))
		copy(glyph, lines)
		chars[r] = glyph
		if lower := toLowerRune(r); lower != r {
			chars[lower] = glyph
		}
	}
	return chars
}

func toLowerRune(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
