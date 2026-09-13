// Package cli es un adaptador primario (driving adapter): traduce
// argumentos de línea de comandos en llamadas al puerto
// ports.BannerGenerator. Es el equivalente Go del "Uso desde frontend:
// await invoke('generate_banner', ...)" del spec original, y deja el
// camino preparado para la "CLI futura" que pedía el documento (acá ya
// es la interfaz principal, dado que el target es Go y no Tauri).
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gabriel/ascii-banner-studio/internal/core/domain"
	"github.com/gabriel/ascii-banner-studio/internal/core/ports"
)

// CLI agrupa las dependencias necesarias para correr el comando.
type CLI struct {
	Generator ports.BannerGenerator
	Exporters ports.ExporterRepository
	Stdout    io.Writer
	Stderr    io.Writer
}

// New crea una CLI lista para usar, con stdout/stderr reales.
func New(generator ports.BannerGenerator, exporters ports.ExporterRepository) *CLI {
	return &CLI{Generator: generator, Exporters: exporters, Stdout: os.Stdout, Stderr: os.Stderr}
}

// Run ejecuta la CLI con los argumentos dados (sin el nombre del
// binario), devolviendo el exit code a usar.
//
// Uso:
//
//	asciibanner "TEXTO" [flags]
//	asciibanner list-fonts
//
// Flags:
//
//	-font string       fuente a usar (default "block")
//	-spacing int       espacio entre caracteres (default 1)
//	-align string      left|center|right (default "left")
//	-uppercase         convierte a mayúsculas
//	-trim              recorta espacios al inicio/fin
//	-format string     si se pasa, exporta en ese formato en vez de imprimir texto plano
func (c *CLI) Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(c.Stderr, "uso: asciibanner \"TEXTO\" [flags] | asciibanner list-fonts")
		return 2
	}

	if args[0] == "list-fonts" {
		for _, f := range c.Generator.ListFonts() {
			fmt.Fprintf(c.Stdout, "%s\t%s\n", f.ID, f.Name)
		}
		return 0
	}

	text := args[0]
	fs := flag.NewFlagSet("asciibanner", flag.ContinueOnError)
	fs.SetOutput(c.Stderr)
	font := fs.String("font", "block", "fuente a usar (ver list-fonts: block, big, banner, minimal + 20 estilos figlet)")
	spacing := fs.Int("spacing", 1, "espacio entre caracteres")
	align := fs.String("align", "left", "alineación: left|center|right")
	uppercase := fs.Bool("uppercase", false, "convertir a mayúsculas")
	trim := fs.Bool("trim", false, "recortar espacios al inicio/fin")
	format := fs.String("format", "", "formato de exportación: txt|javascript|typescript|rust|python|json")

	if err := fs.Parse(args[1:]); err != nil {
		return 2
	}

	opts := domain.BannerOptions{
		Font:      *font,
		Spacing:   *spacing,
		Align:     domain.Alignment(strings.ToLower(*align)),
		Uppercase: *uppercase,
		Trim:      *trim,
	}

	output, err := c.Generator.Generate(text, opts)
	if err != nil {
		fmt.Fprintln(c.Stderr, err.Error())
		return 1
	}

	if *format == "" || *format == "txt" {
		fmt.Fprintln(c.Stdout, output)
		return 0
	}

	exporter, ok := c.Exporters.Get(*format)
	if !ok {
		fmt.Fprintf(c.Stderr, "%s: %s\n", domain.ErrUnknownExportFormat.Error(), *format)
		return 1
	}

	lines := strings.Split(output, "\n")
	exported, err := exporter.Export(lines)
	if err != nil {
		fmt.Fprintln(c.Stderr, err.Error())
		return 1
	}
	fmt.Fprint(c.Stdout, exported)
	return 0
}
