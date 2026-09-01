package domain_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

// testFont construye una fuente mínima para aislar el comportamiento de
// Glyph() y Width() sin depender de los adaptadores concretos.
func testFont() domain.Font {
	return domain.Font{
		Height: 2,
		Characters: map[rune]domain.Glyph{
			'A': {"AA", "AA"},
			'?': {"??", "??"},
		},
		Fallback: '?',
	}
}

func TestGlyphReturnsSupportedCharacter(t *testing.T) {
	f := testFont()
	g := f.Glyph('A')
	if len(g) != 2 || g[0] != "AA" || g[1] != "AA" {
		t.Fatalf("unexpected glyph for 'A': %q", g)
	}
}

func TestGlyphFallsBackForUnknownCharacter(t *testing.T) {
	f := testFont()
	g := f.Glyph('~')
	if g[0] != "??" {
		t.Fatalf("expected fallback glyph, got %q", g[0])
	}
}

func TestGlyphReturnsBlankWhenFallbackMissing(t *testing.T) {
	f := domain.Font{
		Height:     2,
		Characters: map[rune]domain.Glyph{'A': {"AA", "AA"}},
		// Fallback: rune 0 no está en Characters → debe devolver filas vacías.
	}
	g := f.Glyph('~')
	if len(g) != 2 {
		t.Fatalf("expected %d blank rows, got %d", f.Height, len(g))
	}
	for _, row := range g {
		if row != "" {
			t.Fatalf("expected blank row, got %q", row)
		}
	}
}

func TestGlyphWidth(t *testing.T) {
	if (domain.Glyph{}).Width() != 0 {
		t.Fatal("expected width 0 for empty glyph")
	}
	g := domain.Glyph{"ab", "cd"}
	if g.Width() != 2 {
		t.Fatalf("expected width 2, got %d", g.Width())
	}
	unicode := domain.Glyph{"█x"}
	if unicode.Width() != 2 {
		t.Fatalf("expected rune-aware width 2, got %d", unicode.Width())
	}
}
