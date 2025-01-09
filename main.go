package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
	"log"
	"main.go/azure"
	"os"
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
			// Fetch logs for the selected container
			fmt.Printf("we presend enter ")
			// Switch to logs view
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

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	p := azure.GetAzureProjects()
	fmt.Println(p.Value)

	for _, tmrf := range p.Value {
		log.Println("NAME:", *tmrf.Name)
	}

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

	// Style the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#7287fd")).
		Bold(false)
	t.SetStyles(s)

	m := model{
		table: t,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
