package domain_test

import (
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
)

func TestErrTextTooLong(t *testing.T) {
	if domain.ErrTextTooLong.Error() != "Text is too long." {
		t.Fatalf("ErrTextTooLong message mismatch: got %q", domain.ErrTextTooLong)
	}
}

func TestErrFontNotFound(t *testing.T) {
	if domain.ErrFontNotFound.Error() != "font not found" {
		t.Fatalf("ErrFontNotFound message mismatch: got %q", domain.ErrFontNotFound)
	}
}

func TestErrUnknownExportFormat(t *testing.T) {
	if domain.ErrUnknownExportFormat.Error() != "unknown export format" {
		t.Fatalf("ErrUnknownExportFormat message mismatch: got %q", domain.ErrUnknownExportFormat)
	}
}

func TestMaxTextLength(t *testing.T) {
	if domain.MaxTextLength != 4096 {
		t.Fatalf("MaxTextLength: got %d, want 4096", domain.MaxTextLength)
	}
}
