package app

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/azure"
	"main.go/internal/tui"
)

// model holds the data we need in the view
type model struct {
	theme tui.Theme
	table table.Model
}

// no inital commands
func (m model) Init() tea.Cmd {
	return nil
}

// keep the update fast https://leg100.github.io/en/posts/building-bubbletea-programs/
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	return m.table.View()
}

// not a good name
func projectsModel() model {
	// this is to slow?!
	p := azure.GetAzureProjects()

	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Id", Width: 40},
	}

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

	m := model{
		table: t,
	}
	return m
}

// todo add cfg app.Config as arg
func Start() error {
	m := projectsModel()

	a := tea.NewProgram(m,
		tea.WithAltScreen())
	// add routines here
	_, err := a.Run()
	return err
}
