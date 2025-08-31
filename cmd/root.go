package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mapture",
    Short: "Mapture - capture and search your keymaps quickly",
    Long:  "Mapture is a CLI tool that parses config files (like Aerospace, Neovim, Tmux) and lets you fuzzy search your keyboard mappings.",
}

// Execute runs the CLI.
// Called from main.go
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
