package parser

import (
	"sort"
	"strings"

	"github.com/thomasmclean993/mapture/internal/model"
)

// sortKeymaps sorts by Source, Mode, Shortcut, Actions for deterministic tests
func sortKeymaps(k []model.Keymap) {
	sort.Slice(k, func(i, j int) bool {
		if k[i].Source != k[j].Source {
			return k[i].Source < k[j].Source
		}
		if k[i].Mode != k[j].Mode {
			return k[i].Mode < k[j].Mode
		}
		if k[i].Shortcut != k[j].Shortcut {
			return k[i].Shortcut < k[j].Shortcut
		}
		ai := strings.Join(k[i].Actions, ",")
		aj := strings.Join(k[j].Actions, ",")
		return ai < aj
	})
}
