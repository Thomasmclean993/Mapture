package parser

import "github.com/thomasmclean993/mapture/internal/model"

// Parser is the interface that all config parsers (Aerospace, Nvim, Tmux, etc.)
// must implement.
type Parser interface {
    Parse(data []byte) ([]model.Keymap, error)
}
