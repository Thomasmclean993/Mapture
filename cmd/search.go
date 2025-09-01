package cmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/thomasmclean993/mapture/internal/parser"
	"github.com/thomasmclean993/mapture/internal/search"
	"github.com/thomasmclean993/mapture/internal/tui"
)

var searchFilePath string
var searchSource string
var useTUI bool

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Fuzzy search your keymaps",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		if len(args) > 0 {
			query = args[0]
		}

		keymaps, err := parser.GetKeymaps(searchSource, searchFilePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if useTUI {
			program := tea.NewProgram(tui.New(keymaps))
			if _, err := program.Run(); err != nil {
				fmt.Println("TUI error:", err)
			}
			return
		}

		if query == "" {
			fmt.Println("Error: must provide a query (unless using --tui)")
			return
		}

		results := search.Search(keymaps, query)
		if len(results) == 0 {
			fmt.Println("No matches found.")
			return
		}

		for _, km := range results {
			fmt.Printf("[%s] (%s) %s -> %s\n",
				km.Source, km.Mode, km.Shortcut,
				strings.Join(km.Actions, ", "),
			)
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&searchFilePath, "file", "f", "", "Path to config file (optional)")
	// ✅ default directly to "all"
	searchCmd.Flags().StringVarP(&searchSource, "source", "s", "all", "Config source (aerospace, nvim, all). Defaults to all.")
	searchCmd.Flags().BoolVarP(&useTUI, "tui", "t", false, "Interactive TUI search mode")
	rootCmd.AddCommand(searchCmd)
}
