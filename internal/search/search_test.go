package search

import (
    "testing"

    "github.com/google/go-cmp/cmp"
    "github.com/thomasmclean993/mapture/internal/model"
)

func TestSearch(t *testing.T) {
    keymaps := []model.Keymap{
        {Source: "aerospace", Mode: "main", Shortcut: "alt-q", Actions: "close"},
        {Source: "aerospace", Mode: "main", Shortcut: "alt-f", Actions: "fullscreen"},
        {Source: "nvim", Mode: "normal", Shortcut: "<leader>ff", Actions: "telescope find_files"},
        {Source: "nvim", Mode: "normal", Shortcut: "<leader>fg", Actions: "telescope find_grep"},
    }

    got := Search(keymaps, "find")
    want := []model.Keymap{
        
        {Source: "nvim", Mode: "normal", Shortcut: "<leader>ff", Actions: "telescope find_files"},
		{Source: "nvim", Mode: "normal", Shortcut: "<leader>fg", Actions: "telescope find_grep"},
    }

    if diff := cmp.Diff(want, got); diff != "" {
        t.Errorf("Search mismatch (-want +got):\n%s", diff)
    }
}
