package search

import (
	"github.com/lithammer/fuzzysearch/fuzzy"

    "github.com/thomasmclean993/mapture/internal/model"
)

func Search(keymaps []model.Keymap, query string) []model.Keymap {
    var result []model.Keymap

    for _, km := range keymaps {
        if fuzzy.MatchNormalized(query, km.Shortcut) ||
           fuzzy.MatchNormalized(query, km.Action) ||
           fuzzy.MatchNormalized(query, km.Mode) {
            result = append(result, km)
        }
    }

    return result
}
