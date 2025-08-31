package tui

import (
    "fmt"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/thomasmclean993/mapture/internal/model"
    "github.com/thomasmclean993/mapture/internal/search"
)

type TUIModel struct {
    input   textinput.Model
    keymaps []model.Keymap
    results []model.Keymap
}

func New(keymaps []model.Keymap) tea.Model {
    ti := textinput.New()
    ti.Placeholder = "Type to search..."
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 40

    return TUIModel{
        input:   ti,
        keymaps: keymaps,
        results: keymaps,
    }
}

func (m TUIModel) Init() tea.Cmd {
    return textinput.Blink
}

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    m.input, cmd = m.input.Update(msg)

    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" || msg.String() == "esc" {
            return m, tea.Quit
        }
    }

    query := strings.TrimSpace(m.input.Value())
    if query == "" {
        m.results = m.keymaps
    } else {
        m.results = search.Search(m.keymaps, query)
    }

    return m, cmd
}

func (m TUIModel) View() string {
    s := "🔍 Search keymaps: " + m.input.View() + "\n\n"
    for _, km := range m.results {
        s += fmt.Sprintf("[%s] (%s) %s -> %s\n",
            km.Source, km.Mode, km.Shortcut, km.Action)
    }
    s += "\n(esc/ctrl+c to quit)\n"
    return s
}
