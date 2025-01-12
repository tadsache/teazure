package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/tui"
)

type ReposModel struct {
	KeyMap KeyMap // maybe specefiy --

	table table.Model
}

// todo KeyMAp

var reposColumns = []table.Column{
	{Title: "Name", Width: 150},
	// todo
}

func NewReposModel() ReposModel {
	// r := azure.GetRepos
	// todo azure
	data := []string{"ABX", "DEF"}
	var rows []table.Row
	for _, j := range data {
		row := table.Row{
			j,
		}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(reposColumns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	t.SetStyles(tui.TableStyle)

	m := ReposModel{
		table:  t,
		KeyMap: DefaultKeyMap(),
	}
	return m
}

// ELM

func (m ReposModel) Init() tea.Cmd {
	return nil
}

func (m ReposModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Enter):
			m.Enter()
		case key.Matches(msg, m.KeyMap.Quit):
			// fixme not the right way to handle it
			m.Quit()
		}
		// Let the table handle up/down navigation or other keys
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m ReposModel) Enter() {
	fmt.Println("ENTER")
}

func (m ReposModel) Quit() (tea.Model, tea.Cmd) {
	fmt.Println("QUIT")
	return nil, nil
}

func (m ReposModel) View() string {
	return m.table.View()
}
