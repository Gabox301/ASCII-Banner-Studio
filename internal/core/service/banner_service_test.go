package service_test

import (
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
)

func newService() *service.BannerService {
	return service.NewBannerService(fonts.NewRegistry())
}

func baseOpts() domain.BannerOptions {
	return domain.BannerOptions{Font: "block", Spacing: 1, Align: domain.AlignLeft}
}

func TestGeneratesSingleCharacter(t *testing.T) {
	s := newService()
	out, err := s.Generate("A", baseOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 5 {
		t.Fatalf("expected 5 rows for block font, got %d", len(lines))
	}
	if lines[0] == "" {
		t.Fatalf("expected non-empty first row")
	}
}

func TestGeneratesWord(t *testing.T) {
	s := newService()
	out, err := s.Generate("HI", baseOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out == "" {
		t.Fatal("expected non-empty output")
	}
	lines := strings.Split(out, "\n")
	width := len([]rune(lines[0]))
	for _, l := range lines {
		if len([]rune(l)) != width {
			t.Fatalf("rows must have consistent width, got %d vs %d", len([]rune(l)), width)
		}
	}
}

func TestHandlesSpaces(t *testing.T) {
	s := newService()
	out, err := s.Generate("A B", baseOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out == "" {
		t.Fatal("expected non-empty output for text containing spaces")
	}
}

func TestHandlesUnknownCharacter(t *testing.T) {
	s := newService()
	// '~' no está en el charset soportado: no debe hacer panic, debe
	// usar el fallback configurado en la fuente.
	out, err := s.Generate("A~B", baseOpts())
	if err != nil {
		t.Fatalf("unexpected error for unknown character: %v", err)
	}
	if out == "" {
		t.Fatal("expected fallback rendering, got empty output")
	}
}

func TestAppliesSpacing(t *testing.T) {
	s := newService()
	optsNoSpacing := baseOpts()
	optsNoSpacing.Spacing = 0
	optsWithSpacing := baseOpts()
	optsWithSpacing.Spacing = 4

	outNoSpacing, err := s.Generate("AB", optsNoSpacing)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outWithSpacing, err := s.Generate("AB", optsWithSpacing)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	w1 := len([]rune(strings.Split(outNoSpacing, "\n")[0]))
	w2 := len([]rune(strings.Split(outWithSpacing, "\n")[0]))
	if w2 <= w1 {
		t.Fatalf("expected wider output with more spacing: %d vs %d", w2, w1)
	}
}

func TestAppliesAlignment(t *testing.T) {
	s := newService()
	opts := baseOpts()
	opts.Align = domain.AlignRight
	// Dos líneas de ancho distinto para que la alineación sea visible.
	out, err := s.Generate("I\nHI", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	// Todas las filas deben terminar con el mismo ancho total.
	width := len([]rune(lines[0]))
	for _, l := range lines {
		if len([]rune(l)) != width {
			t.Fatalf("aligned rows must share width, got %d vs %d", len([]rune(l)), width)
		}
	}
	// Alineado a la derecha: la primera fila del bloque más angosto
	// debe tener padding de espacios a la izquierda.
	if !strings.HasPrefix(lines[0], " ") {
		t.Fatalf("expected left padding for right alignment, got %q", lines[0])
	}
}

func TestConvertsToUppercase(t *testing.T) {
	s := newService()
	opts := baseOpts()
	opts.Uppercase = true
	lower, err := s.Generate("a", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	upper, err := s.Generate("A", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if lower != upper {
		t.Fatalf("expected uppercase conversion to make 'a' render like 'A'")
	}
}

func TestHandlesEmptyInput(t *testing.T) {
	s := newService()
	out, err := s.Generate("", baseOpts())
	if err != nil {
		t.Fatalf("unexpected error for empty input: %v", err)
	}
	if out != "" {
		t.Fatalf("expected empty output for empty input, got %q", out)
	}
}

func TestRejectsTextTooLong(t *testing.T) {
	s := newService()
	longText := strings.Repeat("A", domain.MaxTextLength+1)
	_, err := s.Generate(longText, baseOpts())
	if err != domain.ErrTextTooLong {
		t.Fatalf("expected ErrTextTooLong, got %v", err)
	}
}

func TestUnknownFontReturnsError(t *testing.T) {
	s := newService()
	opts := baseOpts()
	opts.Font = "does-not-exist"
	_, err := s.Generate("A", opts)
	if err != domain.ErrFontNotFound {
		t.Fatalf("expected ErrFontNotFound, got %v", err)
	}
}

func TestListFontsIncludesMVPFonts(t *testing.T) {
	s := newService()
	fontsList := s.ListFonts()
	if len(fontsList) < 4 {
		t.Fatalf("expected at least 4 fonts, got %d", len(fontsList))
	}
	ids := map[string]bool{}
	for _, f := range fontsList {
		ids[f.ID] = true
	}
	for _, want := range []string{"block", "big", "banner", "minimal"} {
		if !ids[want] {
			t.Fatalf("expected font %q to be registered", want)
		}
	}
}

// minimalOpts usa la fuente minimal (una sola fila, passthrough) con
// spacing 0, lo que permite aserciones exactas de alineación.
func minimalOpts() domain.BannerOptions {
	return domain.BannerOptions{Font: "minimal", Spacing: 0}
}

func TestLeftAlignmentIsDefault(t *testing.T) {
	s := newService()
	out, err := s.Generate("A\nABC", minimalOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "A  " {
		t.Fatalf("expected left padding to the right, got %q", lines[0])
	}
	if lines[2] != "ABC" {
		t.Fatalf("expected widest line unchanged, got %q", lines[2])
	}
}

func TestAppliesCenterAlignment(t *testing.T) {
	s := newService()
	opts := minimalOpts()
	opts.Align = domain.AlignCenter
	out, err := s.Generate("A\nABC", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != " A " {
		t.Fatalf("expected centered ' A ', got %q", lines[0])
	}
}

func TestAppliesRightAlignment(t *testing.T) {
	s := newService()
	opts := minimalOpts()
	opts.Align = domain.AlignRight
	out, err := s.Generate("A\nABC", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "  A" {
		t.Fatalf("expected right-aligned '  A', got %q", lines[0])
	}
}

func TestClampsNegativeSpacing(t *testing.T) {
	s := newService()
	zeroOut, err := s.Generate("AB", minimalOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	negOpts := minimalOpts()
	negOpts.Spacing = -3
	negOut, err := s.Generate("AB", negOpts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if zeroOut != negOut {
		t.Fatalf("expected negative spacing to clamp to 0: %q vs %q", negOut, zeroOut)
	}
}

func TestTrimOptionRemovesPaddingSpaces(t *testing.T) {
	s := newService()
	opts := minimalOpts()
	opts.Trim = true
	out, err := s.Generate("  HI  ", opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "HI" {
		t.Fatalf("expected trimmed output 'HI', got %q", out)
	}
}

func TestHandlesEmptyLineInMultilineText(t *testing.T) {
	s := newService()
	out, err := s.Generate("A\n\nB", minimalOpts())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	// Entrada de 3 bloques ("A", "", "B") → 3 filas + 2 separadores.
	if len(lines) != 5 {
		t.Fatalf("expected 5 output rows, got %d: %q", len(lines), out)
	}
	if lines[0] != "A" || lines[4] != "B" {
		t.Fatalf("unexpected multiline rows: %q", lines)
	}
}

// stubFontRepo es un repositorio de fuentes mínimo para probar el motor con
// fuentes sintéticas (por ejemplo glyphs más cortos que la altura declarada).
type stubFontRepo struct {
	font domain.Font
}

func (s stubFontRepo) Get(id string) (domain.Font, bool) {
	if id == s.font.ID {
		return s.font, true
	}
	return domain.Font{}, false
}

func (s stubFontRepo) List() []domain.FontInfo { return nil }

func TestFillsMissingGlyphRowsWithBlanks(t *testing.T) {
	font := domain.Font{
		ID:         "short-glyph",
		Height:     3,
		Characters: map[rune]domain.Glyph{'A': {"ROW"}},
		Fallback:   '?',
	}
	s := service.NewBannerService(stubFontRepo{font: font})
	out, err := s.Generate("A", domain.BannerOptions{Font: "short-glyph"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 3 {
		t.Fatalf("expected height 3 rows, got %d", len(lines))
	}
	// Las filas faltantes se rellenan en blanco y luego se alinean al ancho
	// máximo del bloque (len("ROW") == 3 → tres espacios).
	if lines[0] != "ROW" || lines[1] != "   " || lines[2] != "   " {
		t.Fatalf("unexpected filled rows: %q", lines)
	}
}
