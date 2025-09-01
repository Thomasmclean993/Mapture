package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thomasmclean993/mapture/internal/model"
)

func TestAerospaceParser(t *testing.T) {
	input := `
    [mode.main.binding]
    alt-q = 'close'
    alt-f = 'fullscreen'
    alt-esc = ['reload-config', 'mode main']
    `

	got, err := AerospaceParser{}.Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []model.Keymap{
		{Source: "aerospace", Mode: "main", Shortcut: "alt-q", Actions: []string{"close"}},
		{Source: "aerospace", Mode: "main", Shortcut: "alt-f", Actions: []string{"fullscreen"}},
		{Source: "aerospace", Mode: "main", Shortcut: "alt-esc", Actions: []string{"reload-config", "mode main"}},
	}

	// Ensure deterministic order before comparing
	sortKeymaps(got)
	sortKeymaps(want)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
