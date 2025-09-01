package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for mapture.
// All subcommands (list, search, etc.) attach under this.
var rootCmd = &cobra.Command{
	Use:   "mapture",
	Short: "Mapture - capture and search keymaps",
	Long: `Mapture is a CLI tool that parses your config files
(Aerospace, Neovim, Tmux, etc.) and lets you fuzzy search
all defined keymaps.`,
}

// Execute runs the root command and handles any error.
// Called by main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
