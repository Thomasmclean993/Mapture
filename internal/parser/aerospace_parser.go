package parser

import (
	"github.com/BurntSushi/toml"
	"github.com/thomasmclean993/mapture/internal/model"
)

// aerospaceConfig models just the pieces of the Aerospace config we care about.
type aerospaceConfig struct {
	Mode map[string]modeConfig `toml:"mode"`
}

type modeConfig struct {
	Binding map[string]interface{} `toml:"binding"`
}

// AerospaceParser parses Aerospace TOML config files.
type AerospaceParser struct{}

func (AerospaceParser) Parse(data []byte) ([]model.Keymap, error) {
	var cfg aerospaceConfig
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	var keymaps []model.Keymap

	for modeName, mode := range cfg.Mode {
		for shortcut, actionVal := range mode.Binding {
			switch v := actionVal.(type) {
			case string:
				// Single action → wrap in slice
				keymaps = append(keymaps, model.Keymap{
					Source:   "aerospace",
					Mode:     modeName,
					Shortcut: shortcut,
					Actions:  []string{v},
				})

			case []interface{}:
				// Multiple actions
				var actions []string
				for _, item := range v {
					if s, ok := item.(string); ok {
						actions = append(actions, s)
					}
				}
				keymaps = append(keymaps, model.Keymap{
					Source:   "aerospace",
					Mode:     modeName,
					Shortcut: shortcut,
					Actions:  actions,
				})
			}
		}
	}

	return keymaps, nil
}
