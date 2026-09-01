// Comando asciibanner: punto de entrada que ensambla el hexágono
// (composition root). Es el único lugar del proyecto que conoce tanto
// el core como todos los adaptadores concretos.
package main

import (
	"os"

	"github.com/gabriel/ascii-banner-studio/internal/adapters/cli"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/export"
	"github.com/gabriel/ascii-banner-studio/internal/adapters/fonts"
	"github.com/gabriel/ascii-banner-studio/internal/core/service"
)

// run ensambla la composition root del comando (idéntica a la de la app de
// escritorio en cuanto a hexágono) y ejecuta la CLI. Está separada de main()
// para poder testear el cableado real en proceso; main() solo traduce el
// exit code en os.Exit.
func run(args []string) int {
	fontRepo := fonts.NewRegistry()
	exportRepo := export.NewRegistry()
	generator := service.NewBannerService(fontRepo)

	app := cli.New(generator, exportRepo)
	return app.Run(args)
}

func main() {
	os.Exit(run(os.Args[1:]))
}
