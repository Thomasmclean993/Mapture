package parser

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/thomasmclean993/mapture/internal/model"
)

type NvimParser struct{}

func (NvimParser) Parse(data []byte) ([]model.Keymap, error) {
	var keymaps []model.Keymap
	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}
		if !strings.HasPrefix(line, "vim.keymap.set") {
			continue
		}

		// Strip prefix/suffix: vim.keymap.set(...)
		inside := strings.TrimPrefix(line, "vim.keymap.set")
		inside = strings.Trim(inside, "()")

		// Split on top-level commas
		parts := splitArgs(inside)
		if len(parts) < 3 {
			continue
		}

		rawMode := strings.TrimSpace(parts[0])
		shortcut := stripOuterQuotes(strings.TrimSpace(parts[1]))
		action := stripOuterQuotes(strings.TrimSpace(parts[2]))

		// Handle multiple modes: {"n","v"} or single "n"
		rawMode = strings.Trim(rawMode, "{} \t")
		modeParts := strings.Split(rawMode, ",")
		for _, m := range modeParts {
			cleanedMode := stripOuterQuotes(strings.TrimSpace(m))
			if cleanedMode == "" {
				continue
			}
			keymaps = append(keymaps, model.Keymap{
				Source:   "nvim",
				Mode:     cleanedMode,
				Shortcut: shortcut,
				Actions:  []string{action},
			})
		}
	}

	return keymaps, scanner.Err()
}

// splitArgs splits on commas at top level (ignores commas inside {})
func splitArgs(s string) []string {
	var args []string
	var current strings.Builder
	depth := 0
	for _, r := range s {
		switch r {
		case '{':
			depth++
			current.WriteRune(r)
		case '}':
			depth--
			current.WriteRune(r)
		case ',':
			if depth == 0 {
				args = append(args, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteRune(r)
			}
		default:
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		args = append(args, strings.TrimSpace(current.String()))
	}
	return args
}

// stripOuterQuotes removes only the outermost quotes if present.
// e.g. "'+y'" -> "+y", "\"<Esc>\"" -> <Esc>
func stripOuterQuotes(s string) string {
	if len(s) >= 2 {
		if (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) ||
			(strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) {
			return s[1 : len(s)-1]
		}
	}
	return s
}
