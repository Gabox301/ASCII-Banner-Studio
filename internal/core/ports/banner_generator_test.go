package ports_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

type mockBannerGenerator struct{}

func (mockBannerGenerator) Generate(text string, opts domain.BannerOptions) (string, error) {
	return text, nil
}

func (mockBannerGenerator) ListFonts() []domain.FontInfo {
	return nil
}

func TestBannerGeneratorInterface(t *testing.T) {
	var _ ports.BannerGenerator = mockBannerGenerator{}
	t.Log("BannerGenerator interface is satisfied by mock implementation")
}
