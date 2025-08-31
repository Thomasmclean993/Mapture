package parser

import "github.com/thomasmclean993/mapture/internal/model"

// Parser is a contract for all config parsers.
// Each parser takes raw config bytes and returns Keymaps.
type Parser interface {
    Parse(data []byte) ([]model.Keymap, error)
}
