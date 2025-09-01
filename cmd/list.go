package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thomasmclean993/mapture/internal/parser"
)

var listFilePath string
var listSource string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List keymaps from configs",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		keymaps, err := parser.GetKeymaps(listSource, listFilePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(keymaps) == 0 {
			fmt.Println("No keymaps found.")
			return
		}

		for _, km := range keymaps {
			fmt.Printf("%s → %s  [%s] (%s)\n",
				km.Shortcut,
				strings.Join(km.Actions, ", "),
				km.Source,
				km.Mode,
	)		}
		},
}

func init() {
	listCmd.Flags().StringVarP(&listFilePath, "file", "f", "", "Path to config file (optional)")
	// ✅ default directly to "all"
	listCmd.Flags().StringVarP(&listSource, "source", "s", "all", "Config source (aerospace, nvim, all). Defaults to all.")
	rootCmd.AddCommand(listCmd)
}
