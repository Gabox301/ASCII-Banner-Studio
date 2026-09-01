package cli_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/cli"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
)

// newTestCLI arma la CLI con los registries reales del hexágono y buffers
// para capturar stdout/stderr.
func newTestCLI() (*cli.CLI, *bytes.Buffer, *bytes.Buffer) {
	generator := service.NewBannerService(fonts.NewRegistry())
	var stdout, stderr bytes.Buffer
	c := cli.New(generator, export.NewRegistry())
	c.Stdout = &stdout
	c.Stderr = &stderr
	return c, &stdout, &stderr
}

func TestNewUsesRealStdio(t *testing.T) {
	generator := service.NewBannerService(fonts.NewRegistry())
	c := cli.New(generator, export.NewRegistry())
	if c.Stdout != os.Stdout {
		t.Fatal("New debería configurar Stdout = os.Stdout")
	}
	if c.Stderr != os.Stderr {
		t.Fatal("New debería configurar Stderr = os.Stderr")
	}
}

func TestRunWithoutArgsShowsUsage(t *testing.T) {
	c, _, stderr := newTestCLI()
	if code := c.Run(nil); code != 2 {
		t.Fatalf("expected exit code 2, got %d", code)
	}
	if !strings.Contains(stderr.String(), "uso:") {
		t.Fatalf("expected usage message, got %q", stderr.String())
	}
}

func TestRunListFonts(t *testing.T) {
	c, stdout, _ := newTestCLI()
	if code := c.Run([]string{"list-fonts"}); code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	for _, id := range []string{"block", "big", "banner", "minimal"} {
		if !strings.Contains(stdout.String(), id) {
			t.Fatalf("expected font %q in listing, got %q", id, stdout.String())
		}
	}
}

func TestRunGeneratesBanner(t *testing.T) {
	c, stdout, _ := newTestCLI()
	if code := c.Run([]string{"HI", "-font", "minimal", "-spacing", "0"}); code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "HI") {
		t.Fatalf("expected banner in output, got %q", stdout.String())
	}
}

func TestRunUnknownFontFails(t *testing.T) {
	c, _, stderr := newTestCLI()
	if code := c.Run([]string{"HI", "-font", "nope"}); code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "font not found") {
		t.Fatalf("expected error on stderr, got %q", stderr.String())
	}
}

func TestRunInvalidFlagFails(t *testing.T) {
	c, _, _ := newTestCLI()
	if code := c.Run([]string{"HI", "-bogus"}); code != 2 {
		t.Fatalf("expected exit code 2, got %d", code)
	}
}

func TestRunExportsJSON(t *testing.T) {
	c, stdout, _ := newTestCLI()
	if code := c.Run([]string{"HI", "-font", "minimal", "-spacing", "0", "-format", "json"}); code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.HasPrefix(stdout.String(), "[") {
		t.Fatalf("expected JSON array, got %q", stdout.String())
	}
}

func TestRunExportsTxt(t *testing.T) {
	c, stdout, _ := newTestCLI()
	if code := c.Run([]string{"HI", "-font", "minimal", "-spacing", "0", "-format", "txt"}); code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "HI") {
		t.Fatalf("expected txt output, got %q", stdout.String())
	}
}

func TestRunUnknownFormatFails(t *testing.T) {
	c, _, stderr := newTestCLI()
	if code := c.Run([]string{"HI", "-font", "minimal", "-format", "nope"}); code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "unknown export format") {
		t.Fatalf("expected error on stderr, got %q", stderr.String())
	}
}
