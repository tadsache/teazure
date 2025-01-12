package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/azure"
	"main.go/internal/tui"
)

type ProjectModel struct {
	KeyMap KeyMap
	//Help   help.Model

	table table.Model
}

type KeyMap struct {
	Enter key.Binding
	Quit  key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Enter: key.NewBinding(
			key.WithKeys("enter", "x"),
			key.WithHelp("x", "enter key"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

var projectColumns = []table.Column{
	{Title: "Name", Width: 20},
	{Title: "Id", Width: 40},
}

func NewAzProjectsTable() ProjectModel {
	// this is to slow?!
	p := azure.GetAzureProjects()
	columns := projectColumns

	var rows []table.Row
	for _, j := range p.Value {
		row := table.Row{
			*j.Name,
			j.Id.String(),
		}
		rows = append(rows, row)
	}

	// abstract the table?
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	t.SetStyles(tui.TableStyle)

	m := ProjectModel{
		table:  t,
		KeyMap: DefaultKeyMap(),
	}
	return m
}

// ELM Architecture

func (m ProjectModel) Init() tea.Cmd {
	return nil
}

func (m ProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// todo

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Enter):
			m.Enter()
		case key.Matches(msg, m.KeyMap.Quit):
			m.Quit()
		}

		// Let the table handle up/down navigation or other keys
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
	// Default fallback
	return m, nil
}

func (m ProjectModel) View() string {
	return m.table.View()
}

func (m *ProjectModel) Enter() {
	fmt.Printf("pressed ENTEr")
}

func (m *ProjectModel) Quit() {
	fmt.Printf("pressed QUIT")
}
