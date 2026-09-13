package fonts_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

func allFonts() []domain.Font {
	return []domain.Font{
		fonts.NewBlockFont(),
		fonts.NewBigFont(),
		fonts.NewBannerFont(),
		fonts.NewMinimalFont(),
		fonts.NewStandardFont(),
		fonts.NewBigFigletFont(),
		fonts.NewSlantFont(),
		fonts.NewBlockFigletFont(),
		fonts.NewShadowFont(),
		fonts.NewBubbleFont(),
		fonts.NewDigitalFont(),
		fonts.NewStarwarsFont(),
		fonts.NewDoomFont(),
		fonts.NewScriptFont(),
		fonts.NewBanner3DFont(),
		fonts.NewIsometricFont(),
		fonts.NewLarry3DFont(),
		fonts.NewOgreFont(),
		fonts.NewGraffitiFont(),
		fonts.NewAnsiShadowFont(),
		fonts.NewColossalFont(),
		fonts.NewPyramidFont(),
		fonts.NewTinkerToyFont(),
		fonts.NewEpicFont(),
	}
}

// TestGlyphHeightConsistency verifica que cada glyph tenga exactamente
// font.Height filas, y que todas las filas de un mismo glyph tengan el
// mismo ancho (invariante requerida por el servicio de generación).
func TestGlyphHeightConsistency(t *testing.T) {
	for _, f := range allFonts() {
		for r, glyph := range f.Characters {
			if len(glyph) != f.Height {
				t.Fatalf("font %s: glyph %q has %d rows, want %d", f.ID, r, len(glyph), f.Height)
			}
			width := glyph.Width()
			for i, row := range glyph {
				if len([]rune(row)) != width {
					t.Fatalf("font %s: glyph %q row %d has inconsistent width", f.ID, r, i)
				}
			}
		}
	}
}

// TestUnknownCharacterFallsBackWithoutPanic reproduce el requisito del
// spec: caracteres no soportados nunca deben provocar panic.
func TestUnknownCharacterFallsBackWithoutPanic(t *testing.T) {
	for _, f := range allFonts() {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("font %s panicked on unsupported character: %v", f.ID, r)
				}
			}()
			glyph := f.Glyph('世') // carácter fuera del charset soportado
			if len(glyph) != f.Height {
				t.Fatalf("font %s: fallback glyph has wrong height", f.ID)
			}
		}()
	}
}

func TestRegistryListAndGet(t *testing.T) {
	reg := fonts.NewRegistry()
	list := reg.List()
	if len(list) != 24 {
		t.Fatalf("expected 24 fonts registered, got %d", len(list))
	}
	if _, ok := reg.Get("block"); !ok {
		t.Fatal("expected block font to be registered")
	}
	if _, ok := reg.Get("nonexistent"); ok {
		t.Fatal("expected nonexistent font lookup to fail")
	}
	// Sin duplicaciones: IDs y nombres únicos.
	seenIDs := map[string]bool{}
	seenNames := map[string]bool{}
	for _, info := range list {
		if seenIDs[info.ID] {
			t.Fatalf("duplicate font ID %q", info.ID)
		}
		seenIDs[info.ID] = true
		if seenNames[info.Name] {
			t.Fatalf("duplicate font Name %q", info.Name)
		}
		seenNames[info.Name] = true
	}
	// Las FIGlet que colisionaban con las base usan IDs propios.
	for _, id := range []string{"big-figlet", "block-figlet", "banner3d"} {
		if _, ok := reg.Get(id); !ok {
			t.Fatalf("expected figlet font %q to be registered", id)
		}
	}
}

func TestBigFontIsScaledBlockFont(t *testing.T) {
	block := fonts.NewBlockFont()
	big := fonts.NewBigFont()
	if big.Height != block.Height*2 {
		t.Fatalf("expected big font height to double block font height: %d vs %d", big.Height, block.Height)
	}
}
