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

// escapeKotlin escapa para strings "de Kotlin: además de backslash y
// comillas dobles, el signo $ inicia templates y debe neutralizarse.
func escapeKotlin(s string) string {
	s = escapeDoubleQuoted(s)
	s = strings.ReplaceAll(s, `$`, `${'$'}`)
	return s
}

// escapeRubyDoubleQuoted escapa para strings "de Ruby: #{ inicia
// interpolación y debe neutralizarse (el backslash ya escapado lo cubre
// para \( pero no para #).
func escapeRubyDoubleQuoted(s string) string {
	s = escapeDoubleQuoted(s)
	s = strings.ReplaceAll(s, `#{`, `\#{`)
	return s
}

// escapeDartSingleQuoted escapa para strings 'de Dart: $ inicia
// interpolación ($x o ${...}) y debe neutralizarse.
func escapeDartSingleQuoted(s string) string {
	s = escapeSingleQuoted(s)
	s = strings.ReplaceAll(s, `$`, `\$`)
	return s
}

// escapeShellSingleQuoted escapa para strings 'de Shell: dentro de comillas
// simples no hay escapes, la comilla se cierra, se agrega \' y se reabre.
func escapeShellSingleQuoted(s string) string {
	return strings.ReplaceAll(s, `'`, `'\''`)
}

// escapePowerShellSingleQuoted escapa para strings 'de PowerShell: la
// comilla simple se duplica.
func escapePowerShellSingleQuoted(s string) string {
	return strings.ReplaceAll(s, `'`, `''`)
}
