package parser

import (
	"os"
	"path/filepath"

	"github.com/thomasmclean993/mapture/internal/model"
)

// Registry maps supported config sources to their parser implementations.
var Registry = map[string]Parser{
	"aerospace": AerospaceParser{},
	"nvim":      NvimParser{},
}

// GetKeymaps resolves keymaps for a given source + file path.
// If source == "all", it runs every registered parser and merges results.
func GetKeymaps(source string, filePath string) ([]model.Keymap, error) {
	if source == "all" {
		var all []model.Keymap
		for name, p := range Registry {
			path := filePathFor(name, filePath)
			if path == "" {
				continue
			}
			data, err := os.ReadFile(path)
			if err != nil {
				continue // skip missing file
			}
			keymaps, err := p.Parse(data)
			if err == nil {
				all = append(all, keymaps...)
			}
		}
		return all, nil
	}

	// Single-source path
	p, ok := Registry[source]
	if !ok {
		return nil, ErrUnknownSource(source)
	}

	path := filePathFor(source, filePath)
	if path == "" {
		return nil, ErrNoDefaultPath(source)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.Parse(data)
}

// ErrUnknownSource is returned when no parser is found.
type ErrUnknownSource string

func (e ErrUnknownSource) Error() string {
	return "unknown source: " + string(e)
}

// ErrNoDefaultPath is returned when no default path is available.
type ErrNoDefaultPath string

func (e ErrNoDefaultPath) Error() string {
	return "no default config path for source: " + string(e)
}

// filePathFor picks the correct file path for the given source.
// If the user provided --file, use that. Otherwise fall back to standard defaults.
func filePathFor(source, provided string) string {
	if provided != "" {
		return provided
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "" // no home dir; bail
	}

	switch source {
	case "aerospace":
		return filepath.Join(home, ".config", "aerospace.toml")
	case "nvim":
		return filepath.Join(home, ".config", "nvim", "init.lua")
	// TODO: add more parsers (tmux, i3, etc.)
	default:
		return ""
	}
}
