package gui

import (
	"context"
	"os"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func newTestHandler() *Handler {
	return New(
		service.NewBannerService(fonts.NewRegistry()),
		export.NewRegistry(),
	)
}

func TestSaveBannerDialogError(t *testing.T) {
	orig := saveFileDialog
	defer func() { saveFileDialog = orig }()
	saveFileDialog = func(ctx context.Context, opts runtime.SaveDialogOptions) (string, error) {
		return "", context.DeadlineExceeded
	}
	h := newTestHandler()
	h.SetContext(context.Background())
	_, err := h.SaveBanner("content", "banner.txt")
	if err != context.DeadlineExceeded {
		t.Fatalf("expected dialog error, got %v", err)
	}
}

func TestSaveBannerUserCancels(t *testing.T) {
	orig := saveFileDialog
	defer func() { saveFileDialog = orig }()
	saveFileDialog = func(ctx context.Context, opts runtime.SaveDialogOptions) (string, error) {
		return "", nil
	}
	h := newTestHandler()
	h.SetContext(context.Background())
	path, err := h.SaveBanner("content", "banner.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path != "" {
		t.Fatalf("expected empty path on cancel, got %q", path)
	}
}

func TestSaveBannerWritesFile(t *testing.T) {
	orig := saveFileDialog
	defer func() { saveFileDialog = orig }()
	tmpDir := t.TempDir()
	want := "hello banner"
	saveFileDialog = func(ctx context.Context, opts runtime.SaveDialogOptions) (string, error) {
		return tmpDir + string(os.PathSeparator) + "out.txt", nil
	}
	h := newTestHandler()
	h.SetContext(context.Background())
	path, err := h.SaveBanner(want, "out.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if path != tmpDir+string(os.PathSeparator)+"out.txt" {
		t.Fatalf("unexpected path: %q", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read saved file: %v", err)
	}
	if string(data) != want {
		t.Fatalf("file content mismatch: got %q, want %q", string(data), want)
	}
}

func TestSaveBannerWriteFileError(t *testing.T) {
	orig := saveFileDialog
	defer func() { saveFileDialog = orig }()
	saveFileDialog = func(ctx context.Context, opts runtime.SaveDialogOptions) (string, error) {
		return `\invalid\path\out.txt`, nil
	}
	h := newTestHandler()
	h.SetContext(context.Background())
	_, err := h.SaveBanner("content", "out.txt")
	if err == nil {
		t.Fatal("expected error writing to invalid path")
	}
}
