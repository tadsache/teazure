package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/azure"
	"main.go/internal/tui"
)

type ProjectModel struct {
	table table.Model
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
		table: t,
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
		switch msg.String() {

		// todo make a KEYMAP for the cases
		case "ctrl+c", "q":
			// Quit from table view
			return m, tea.Quit

		case "enter":
			fmt.Printf("we presend enter ")
			return m, nil
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
