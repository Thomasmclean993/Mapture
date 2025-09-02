package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/thomasmclean993/mapture/internal/config"
)

type SourceConfig struct {
	Path string `yaml:"path"`
}
type Sources struct {
	Sources map[string]SourceConfig `yaml:"sources"`
}

var configFilePath string
var source string
var path string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Mapture configuration",
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or update a source filepath in sources.yml",
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" || path == "" {
			fmt.Println("Error: --source and --path are required")
			return
		}

		if configFilePath == "" {
			home, _ := os.UserHomeDir()
			configFilePath = filepath.Join(home, ".config", "mapture", "sources.yml")
		}

		srcs := Sources{Sources: make(map[string]SourceConfig)}

		data, err := os.ReadFile(configFilePath)
		if err == nil {
			yaml.Unmarshal(data, &srcs)
		}

		srcs.Sources[source] = SourceConfig{Path: path}

		os.MkdirAll(filepath.Dir(configFilePath), 0755)

		out, _ := yaml.Marshal(&srcs)
		if err := os.WriteFile(configFilePath, out, 0644); err != nil {
			fmt.Println("Failed to write config file:", err)
			return
		}

		fmt.Printf("Updated %s → %s in %s\n", source, path, configFilePath)
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List resolved sources and their config filepaths",
	Run: func(cmd *cobra.Command, args []string) {
		if configFilePath == "" {
			home, _ := os.UserHomeDir()
			configFilePath = filepath.Join(home, ".config", "mapture", "sources.yml")
		}

		srcs, err := config.LoadSources(configFilePath)
		if err != nil {
			fmt.Println("Error loading sources:", err)
			return
		}

		fmt.Println("📂 Resolved config sources:")
		for name, sc := range srcs.Sources {
			fmt.Printf("  - %-10s → %s\n", name, sc.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Add subcommands
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configListCmd)

	// Flags for `add`
	configAddCmd.Flags().StringVar(&source, "source", "", "Source name (aerospace, nvim, tmux, etc.)")
	configAddCmd.Flags().StringVar(&path, "path", "", "Custom filepath for this source")
	configAddCmd.Flags().StringVar(&configFilePath, "file", "", "Optional path to sources.yml")

	// Flag for `list`
	configListCmd.Flags().StringVar(&configFilePath, "file", "", "Optional path to sources.yml")
}
