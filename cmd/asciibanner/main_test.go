package main

import "testing"

// run() es la composition root del binario: usa los registries reales y
// delega en la CLI. Estos tests cubren la integración completa (códigos de
// salida reales), no solo los adaptadores por separado.
func TestRunCompositionRoot(t *testing.T) {
	if code := run([]string{"list-fonts"}); code != 0 {
		t.Fatalf("list-fonts: expected exit 0, got %d", code)
	}
	if code := run([]string{"HI", "-font", "minimal", "-spacing", "0"}); code != 0 {
		t.Fatalf("generar banner: expected exit 0, got %d", code)
	}
	if code := run([]string{"HI", "-font", "nope"}); code != 1 {
		t.Fatalf("fuente inexistente: expected exit 1, got %d", code)
	}
	if code := run(nil); code != 2 {
		t.Fatalf("sin argumentos: expected exit 2, got %d", code)
	}
}
