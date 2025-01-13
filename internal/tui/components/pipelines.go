package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/tui"
)

type PipelinesModel struct {
	KeyMap KeyMap

	table table.Model
}

// fixme KeyMap for all tableViews can be the same ...

var pipelineColumns = []table.Column{
	{Title: "Name", Width: 50},
}

func NewPipelinesModel() *PipelinesModel {

	var test = []string{"erste line", "zweite line"}

	var rows []table.Row
	for _, r := range test {
		row := table.Row{
			r,
		}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(pipelineColumns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	t.SetStyles(tui.TableStyle)

	m := &PipelinesModel{
		table:  t,
		KeyMap: DefaultKeyMap(), // fixme
	}
	return m
}

// ELM

func (m PipelinesModel) Init() tea.Cmd {
	return nil
}

func (m PipelinesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Enter):
			m.Enter()
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit
		}

		// Delegate message handling to the table and update its state
		var cmd tea.Cmd
		updatedTable, tableCmd := m.table.Update(msg)
		m.table = updatedTable
		cmd = tea.Batch(cmd, tableCmd)

		return m, cmd
	}
	// Default fallback
	return m, nil
}

func (m PipelinesModel) View() string {
	return m.table.View()
}

func (m PipelinesModel) Enter() {
	fmt.Println("ENTER")
}
