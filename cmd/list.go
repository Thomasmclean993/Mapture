package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/thomasmclean993/mapture/internal/parser"
)

// flags
var listFilePath string
var listSource string

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List keymaps from a config",
    Args:  cobra.NoArgs, // no positional args now, since we use --source
    Run: func(cmd *cobra.Command, args []string) {
        p, ok := parser.Registry[listSource]
        if !ok {
            fmt.Printf("Unknown source: %s\n", listSource)
            return
        }

        data, err := os.ReadFile(listFilePath)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        keymaps, err := p.Parse(data)
        if err != nil {
            fmt.Println("Parse error:", err)
            return
        }

        for _, km := range keymaps {
            fmt.Printf("[%s] (%s) %s -> %s\n",
                km.Source, km.Mode, km.Shortcut, km.Actions)
        }
    },
}

func init() {
    // --file flag
    listCmd.Flags().StringVarP(
        &listFilePath,
        "file", "f",
        "aerospace.toml",
        "Path to config file",
    )

    // --source flag (aerospace by default)
    listCmd.Flags().StringVarP(
        &listSource,
        "source", "s",
        "aerospace",
        "Config source (aerospace, nvim, tmux)",
    )

    rootCmd.AddCommand(listCmd)
}
