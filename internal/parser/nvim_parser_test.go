package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thomasmclean993/mapture/internal/model"
)

func TestNvimParser(t *testing.T) {
	input := `
    vim.keymap.set("n", "<leader>ff", ":Telescope find_files<CR>")
    vim.keymap.set("i", "jk", "<Esc>")
    vim.keymap.set({"n","v"}, "<leader>y", '"+y')
    `

	got, err := NvimParser{}.Parse([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []model.Keymap{
		{Source: "nvim", Mode: "n", Shortcut: "<leader>ff", Actions: []string{":Telescope find_files<CR>"}},
		{Source: "nvim", Mode: "i", Shortcut: "jk", Actions: []string{"<Esc>"}},
		{Source: "nvim", Mode: "n", Shortcut: "<leader>y", Actions: []string{`"+y`}},
		{Source: "nvim", Mode: "v", Shortcut: "<leader>y", Actions: []string{`"+y`}},
	}

	// Ensure deterministic order before comparing
	sortKeymaps(got)
	sortKeymaps(want)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
