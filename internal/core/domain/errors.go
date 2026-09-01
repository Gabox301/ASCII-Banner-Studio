package domain

import "errors"

// ErrTextTooLong se devuelve cuando la entrada excede MaxTextLength.
var ErrTextTooLong = errors.New("Text is too long.")

// ErrFontNotFound se devuelve cuando se solicita una fuente inexistente.
var ErrFontNotFound = errors.New("font not found")

// ErrUnknownExportFormat se devuelve cuando se pide un formato de
// exportación no registrado.
var ErrUnknownExportFormat = errors.New("unknown export format")

// MaxTextLength es la longitud máxima de entrada soportada en el MVP.
const MaxTextLength = 4096
