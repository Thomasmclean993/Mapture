package parser

import (
    "github.com/BurntSushi/toml"
    "github.com/thomasmclean993/mapture/internal/model"
)

// aerospaceConfig models just the pieces of the aerospace config we care about.
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

    for modeName, mode := range cfg.Mode { // e.g. "main"
        for shortcut, actionVal := range mode.Binding {
            switch v := actionVal.(type) {
            case string:
                keymaps = append(keymaps, model.Keymap{
                    Source:   "aerospace",
                    Mode:     modeName, // NEW: preserve mode!
                    Shortcut: shortcut,
                    Action:   v,
                })
            case []interface{}:
                for _, item := range v {
                    if s, ok := item.(string); ok {
                        keymaps = append(keymaps, model.Keymap{
                            Source:   "aerospace",
                            Mode:     modeName, // NEW
                            Shortcut: shortcut,
                            Action:   s,
                        })
                    }
                }
            }
        }
    }

    return keymaps, nil
}

