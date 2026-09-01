package domain_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

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

func TestAlignmentConstants(t *testing.T) {
	if domain.AlignLeft != "left" {
		t.Fatalf("AlignLeft: got %q, want %q", domain.AlignLeft, "left")
	}
	if domain.AlignCenter != "center" {
		t.Fatalf("AlignCenter: got %q, want %q", domain.AlignCenter, "center")
	}
	if domain.AlignRight != "right" {
		t.Fatalf("AlignRight: got %q, want %q", domain.AlignRight, "right")
	}
}

func TestBannerOptions(t *testing.T) {
	opts := domain.BannerOptions{
		Font:      "big",
		Spacing:   2,
		Align:     domain.AlignCenter,
		Uppercase: true,
		Trim:      false,
	}
	if opts.Font != "big" {
		t.Fatalf("Font: got %q, want %q", opts.Font, "big")
	}
	if opts.Spacing != 2 {
		t.Fatalf("Spacing: got %d, want 2", opts.Spacing)
	}
	if opts.Align != domain.AlignCenter {
		t.Fatalf("Align: got %q, want %q", opts.Align, domain.AlignCenter)
	}
	if !opts.Uppercase {
		t.Fatal("Uppercase: got false, want true")
	}
	if opts.Trim {
		t.Fatal("Trim: got true, want false")
	}
}

func TestFontStruct(t *testing.T) {
	f := domain.Font{
		ID:     "big",
		Name:   "Big",
		Height: 5,
		Characters: map[rune]domain.Glyph{
			'A': {"  A  ", " A A ", "AAAAA", "A   A", "A   A"},
		},
		Fallback: '?',
	}
	if f.ID != "big" {
		t.Fatalf("ID: got %q, want %q", f.ID, "big")
	}
	if f.Name != "Big" {
		t.Fatalf("Name: got %q, want %q", f.Name, "Big")
	}
	if f.Height != 5 {
		t.Fatalf("Height: got %d, want 5", f.Height)
	}
	if len(f.Characters) != 1 {
		t.Fatalf("Characters: got %d entries, want 1", len(f.Characters))
	}
	if f.Fallback != '?' {
		t.Fatalf("Fallback: got %q, want %q", f.Fallback, '?')
	}
}

func TestFontInfoStruct(t *testing.T) {
	info := domain.FontInfo{ID: "big", Name: "Big"}
	if info.ID != "big" {
		t.Fatalf("ID: got %q, want %q", info.ID, "big")
	}
	if info.Name != "Big" {
		t.Fatalf("Name: got %q, want %q", info.Name, "Big")
	}
}

func TestFontGlyphReturnsBlankWhenFallbackMissing(t *testing.T) {
	f := domain.Font{
		Height:     3,
		Characters: map[rune]domain.Glyph{'A': {"A", "A", "A"}},
	}
	g := f.Glyph('~')
	if len(g) != 3 {
		t.Fatalf("expected %d blank rows, got %d", f.Height, len(g))
	}
	for i, row := range g {
		if row != "" {
			t.Fatalf("row %d: expected blank, got %q", i, row)
		}
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

func TestGlyphWidthEmpty(t *testing.T) {
	if (domain.Glyph{}).Width() != 0 {
		t.Fatal("expected width 0 for empty glyph")
	}
}

func TestGlyphWidthMethod(t *testing.T) {
	g := domain.Glyph{"abc", "def"}
	if g.Width() != 3 {
		t.Fatalf("expected width 3, got %d", g.Width())
	}
}

func TestGlyphWidthMethodUnicode(t *testing.T) {
	g := domain.Glyph{"█x"}
	if g.Width() != 2 {
		t.Fatalf("expected width 2, got %d", g.Width())
	}
}
