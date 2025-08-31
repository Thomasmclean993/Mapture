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

    parser := AerospaceParser{}
    got, err := parser.Parse([]byte(input))
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

   want := []model.Keymap{
    {Source: "aerospace", Mode: "main", Shortcut: "alt-q", Action: "close"},
    {Source: "aerospace", Mode: "main", Shortcut: "alt-f", Action: "fullscreen"},
    {Source: "aerospace", Mode: "main", Shortcut: "alt-esc", Action: "reload-config"},
    {Source: "aerospace", Mode: "main", Shortcut: "alt-esc", Action: "mode main"},
}
    if diff := cmp.Diff(want, got); diff != "" {
    t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
