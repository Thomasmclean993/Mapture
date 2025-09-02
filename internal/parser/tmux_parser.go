package parser

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/thomasmclean993/mapture/internal/model"
)

// TmuxParser parses tmux key bindings (bind-key lines).
type TmuxParser struct{}

func (TmuxParser) Parse(data []byte) ([]model.Keymap, error) {
	var keymaps []model.Keymap
	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip comments and empty lines
		}

		// Only care about bind-key lines
		if strings.HasPrefix(line, "bind-key") {
			// Remove "bind-key" prefix
			parts := strings.Fields(line)
			if len(parts) < 3 {
				continue
			}

			// parts[0] = bind-key
			// parts[1] = flag or shortcut
			// Without -n:
			//   bind-key C-b split-window -h
			// With -n:
			//   bind-key -n C-p next-window
			var shortcut string
			var actionParts []string

			if parts[1] == "-n" && len(parts) >= 4 {
				shortcut = parts[2]
				actionParts = parts[3:]
			} else {
				shortcut = parts[1]
				actionParts = parts[2:]
			}

			action := strings.Join(actionParts, " ")

			keymaps = append(keymaps, model.Keymap{
				Source:   "tmux",
				Mode:     "global",
				Shortcut: shortcut,
				Actions:  []string{action},
			})
		}
	}

	return keymaps, scanner.Err()
}
