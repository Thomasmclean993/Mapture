package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	m "github.com/thomasmclean993/mapture/internal/model"
	"github.com/thomasmclean993/mapture/internal/search"
)

type TUIModel struct {
	input   textinput.Model
	keymaps []m.Keymap
	results []m.Keymap
	table   table.Model
}

// Create new TUIModel with initial keymaps
func New(keymaps []m.Keymap) tea.Model {
	ti := textinput.New()
	ti.Placeholder = "Type to search..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	// initial empty table
	columns := []table.Column{
		{Title: "Shortcut", Width: 20},
		{Title: "Actions", Width: 40},
		{Title: "Source", Width: 12},
		{Title: "Mode", Width: 12},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithHeight(15),
	)

	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	style.Selected = style.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(style)

	return TUIModel{
		input:   ti,
		keymaps: keymaps,
		results: keymaps,
		table:   t,
	}
}

func (m TUIModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.table, _ = m.table.Update(msg)

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
	m.refreshTable()

	return m, cmd
}

func (m *TUIModel) refreshTable() {
	var rows []table.Row
	for _, km := range m.results {
		rows = append(rows, table.Row{
			km.Shortcut,
			strings.Join(km.Actions, ", "),
			km.Source,
			km.Mode,
		})
	}
	m.table.SetRows(rows)
}

func (m TUIModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		"🔑 Mapture TUI Search (type to filter, esc/ctrl+c to quit)\n"+m.input.View()+"\n",
		m.table.View(),
	)
}
