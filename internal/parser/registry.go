package parser

import (
	"os"
	"path/filepath"

	"github.com/thomasmclean993/mapture/internal/config"
	"github.com/thomasmclean993/mapture/internal/model"
)

// Registry maps supported config sources to their parser implementations.
var Registry = map[string]Parser{
	"aerospace": AerospaceParser{},
	"nvim":      NvimParser{},
	"tmux":      TmuxParser{},
}

// GetKeymaps resolves keymaps for a given source.
// If source == "all", it runs every registered parser and merges results.
func GetKeymaps(source string, configPath string) ([]model.Keymap, error) {
	// If configPath is unset, look in ~/.config/mapture/sources.yml
	if configPath == "" {
		home, _ := os.UserHomeDir()
		configPath = filepath.Join(home, ".config", "mapture", "sources.yml")
	}

	srcs, err := config.LoadSources(configPath)
	if err != nil {
		return nil, err
	}

	// If source == "all", merge everything
	if source == "all" {
		var all []model.Keymap
		for name, p := range Registry {
			path := os.ExpandEnv(srcs.Sources[name].Path)
			data, err := os.ReadFile(path)
			if err != nil {
				continue
			}
			keymaps, err := p.Parse(data)
			if err == nil {
				all = append(all, keymaps...)
			}
		}
		return all, nil
	}

	// Single source
	p, ok := Registry[source]
	if !ok {
		return nil, ErrUnknownSource(source)
	}

	path := os.ExpandEnv(srcs.Sources[source].Path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return p.Parse(data)
}

// ErrUnknownSource is returned when a parser source isn’t registered.
type ErrUnknownSource string

func (e ErrUnknownSource) Error() string {
	return "unknown source: " + string(e)
}
