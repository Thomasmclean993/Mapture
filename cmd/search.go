package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/thomasmclean993/mapture/internal/parser"
    "github.com/thomasmclean993/mapture/internal/search"
)

// shared with list.go for consistent flag naming
var searchFilePath string

var searchCmd = &cobra.Command{
    Use:   "search [query]",
    Short: "Fuzzy search your keymaps by query",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        query := args[0]

        // For now, hardcode Aerospace parser
        p := parser.AerospaceParser{}

        data, err := os.ReadFile(searchFilePath)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        keymaps, err := p.Parse(data)
        if err != nil {
            fmt.Println("Parse error:", err)
            return
        }

        results := search.Search(keymaps, query)
        if len(results) == 0 {
            fmt.Println("No matches found.")
            return
        }

        for _, km := range results {
            fmt.Printf("[%s] (%s) %s -> %s\n",
                km.Source, km.Mode, km.Shortcut, km.Action)
        }
    },
}

func init() {
    searchCmd.Flags().StringVarP(
        &searchFilePath,
        "file", "f",
        "aerospace.toml",
        "Path to config file",
    )

    rootCmd.AddCommand(searchCmd)
}
