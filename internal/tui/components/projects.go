package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/azure"
	"main.go/internal/tui"
)

type ProjectsModel struct {
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
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter key"),
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

func NewProjectsModel() ProjectsModel {
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

	m := ProjectsModel{
		table:  t,
		KeyMap: DefaultKeyMap(),
	}
	return m
}

// ELM Architecture

func (m ProjectsModel) Init() tea.Cmd {
	return nil
}

func (m ProjectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// todo

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Enter):
			// ..
			cmd := m.Enter()
			return m, cmd
		case key.Matches(msg, m.KeyMap.Quit):
			// fixme not the right way to handle it
			return m.Quit()
		}

		// Let the table handle up/down navigation or other keys
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
	// Default fallback
	return m, nil
}

func (m ProjectsModel) View() string {
	return m.table.View()
}

func (m *ProjectsModel) Enter() tea.Cmd {
	fmt.Printf("pressed Enter")
	// switch to project detail view

	// The user pressed Enter on a project
	// We’ll handle "switching to detail" in the parent model,
	// so let's send a custom message upward:
	ms := SelectProjectMsg{
		ProjectName: m.table.SelectedRow()[0],
		ProjectId:   m.table.SelectedRow()[1],
	}

	return func() tea.Msg {
		return ms
	}
}

// SelectProjectMsg is a custom message we’ll emit when the user selects a project
type SelectProjectMsg struct {
	ProjectName string
	ProjectId   string
}

func (m *ProjectsModel) Quit() (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
