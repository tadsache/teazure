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
	table table.Model
}

// no inital commands
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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

// todo add cfg app.Config as arg
func Start() error {
	p := azure.GetAzureProjects()
	fmt.Println(p.Value)

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

	a := tea.NewProgram(m,
		tea.WithAltScreen())
	// add routines here
	_, err := a.Run()
	return err
}
