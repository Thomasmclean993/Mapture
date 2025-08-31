package cmd

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/spf13/cobra"

    "github.com/thomasmclean993/mapture/internal/parser"
    "github.com/thomasmclean993/mapture/internal/search"
    "github.com/thomasmclean993/mapture/internal/tui"
)
// shared with list.go for consistent flag naming
var searchFilePath string
var useTUI bool

var searchCmd = &cobra.Command{
    Use:   "search [query]",
    Short: "Fuzzy search your keymaps by query",
    Args:  cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
    var query string 
    if len(args) > 0 {
        query = args[0]
    }

    // Hardcode Aerospace parser (more later)
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

    if useTUI {
        program := tea.NewProgram(tui.New(keymaps))
        if _, err := program.Run(); err != nil {
            fmt.Println("TUI error:", err)
        }
        return
    }

    // If no query is given in non-TUI mode
    if query == "" {
        fmt.Println("Error: you must provide a query (or use --tui)")
        return
    }

    // normal CLI search
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

    searchCmd.Flags().BoolVarP(&useTUI, "tui", "t", false, "Run search in interactive TUI mode")

    rootCmd.AddCommand(searchCmd)
}
