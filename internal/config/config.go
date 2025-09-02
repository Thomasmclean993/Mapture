package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type SourceConfig struct {
	Path string `yaml:"path"`
}

type Sources struct {
	Sources map[string]SourceConfig `yaml:"sources"`
}

// Default paths if none provided in YAML
var defaults = map[string]string{
	"aerospace": filepath.Join(os.Getenv("HOME"), ".config", "aerospace.toml"),
	"nvim":      filepath.Join(os.Getenv("HOME"), ".config", "nvim", "init.lua"),
	"tmux":      filepath.Join(os.Getenv("HOME"), ".tmux.conf"),
}

// LoadSources loads a YAML config from path, returning struct with defaults filled in.
func LoadSources(configPath string) (Sources, error) {
	srcs := Sources{Sources: make(map[string]SourceConfig)}

	// Start with defaults
	for k, v := range defaults {
		srcs.Sources[k] = SourceConfig{Path: v}
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		// No config file? Just return defaults silently
		return srcs, nil
	}

	if err := yaml.Unmarshal(data, &srcs); err != nil {
		return srcs, err
	}

	// Fill in any missing with defaults
	for k, v := range defaults {
		if _, ok := srcs.Sources[k]; !ok {
			srcs.Sources[k] = SourceConfig{Path: v}
		}
	}

	return srcs, nil
}

