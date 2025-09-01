package model

// Keymap is our universal representation of a keybinding.
// It doesn't care what tool it's from – just captures the intent.
type Keymap struct {
    Source   string // e.g. "aerospace", "nvim", "tmux"
	Mode 	 string // e.g. "main", "resize", "manage"
    Shortcut string // e.g. "mod+f", "<leader>ff"
    Actions   []string // e.g. "toggle_fullscreen"
}

// Maybe a subtype?
