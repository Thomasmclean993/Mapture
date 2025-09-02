package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thomasmclean993/mapture/internal/model"
)

func TestTmuxParser(t *testing.T) {
	input := `
# A comment
bind-key -n C-p next-window
bind-key C-b split-window -h
bind-key -n M-1 select-window -t 1
`

	got, err := TmuxParser{}.Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []model.Keymap{
		{Source: "tmux", Mode: "global", Shortcut: "C-p", Actions: []string{"next-window"}},
		{Source: "tmux", Mode: "global", Shortcut: "C-b", Actions: []string{"split-window -h"}},
		{Source: "tmux", Mode: "global", Shortcut: "M-1", Actions: []string{"select-window -t 1"}},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
