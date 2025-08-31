package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/thomasmclean993/mapture/internal/parser"
)

// Global variable where --file flag value will land
var filePath string

var listCmd = &cobra.Command{
    Use:   "list [source]",
    Short: "List keymaps from a config",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        source := args[0]

        var p parser.Parser
        switch source {
        case "aerospace":
            p = parser.AerospaceParser{}
        default:
            fmt.Printf("Unknown source: %s\n", source)
            return
        }

        data, err := os.ReadFile(filePath) // <-- value from --file flag
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
            fmt.Printf("[%s] (%s) %s -> %s\n", km.Source, km.Mode, km.Shortcut, km.Action)
        }
    },
}

func init() {
    // Attach flag to the command so Cobra knows about it
    listCmd.Flags().StringVarP(
        &filePath,          // pointer to store value
        "file", "f",        // name and shorthand
        "aerospace.toml",   // default value
        "Path to config file", // help text
    )

    rootCmd.AddCommand(listCmd)
}
