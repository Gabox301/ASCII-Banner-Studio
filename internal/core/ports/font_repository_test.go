package ports_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

type mockFontRepository struct {
	fonts map[string]domain.Font
}

func (m mockFontRepository) Get(id string) (domain.Font, bool) {
	f, ok := m.fonts[id]
	return f, ok
}

func (m mockFontRepository) List() []domain.FontInfo {
	return nil
}

func TestFontRepositoryInterface(t *testing.T) {
	m := mockFontRepository{fonts: map[string]domain.Font{"big": {ID: "big", Height: 5}}}
	var _ ports.FontRepository = m
	t.Log("FontRepository interface is satisfied by mock implementation")
}
