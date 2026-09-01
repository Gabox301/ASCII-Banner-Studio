// Package export contiene los adaptadores secundarios que transforman
// el banner generado en distintos formatos de código/texto. Cada
// exporter está aislado del renderer (domain/service) y del resto de
// exporters.
package export

import "strings"

// escapeSingleQuoted escapa backslashes y comillas simples, para
// strings 'entre comillas simples' (JavaScript/TypeScript).
func escapeSingleQuoted(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

// escapeDoubleQuoted escapa backslashes y comillas dobles, para
// strings "entre comillas dobles" (Rust, Python, JSON manual).
func escapeDoubleQuoted(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return s
}

// escapeBacktick escapa backticks y ${ (por si en el futuro se exporta
// como template literal de JS/TS).
func escapeBacktick(s string) string {
	s = strings.ReplaceAll(s, "`", "\\`")
	s = strings.ReplaceAll(s, "${", "\\${")
	return s
}
