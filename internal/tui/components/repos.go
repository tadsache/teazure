package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/azure"
	"main.go/internal/tui"
)

type ReposModel struct {
	KeyMap KeyMap // maybe specefiy --

	table table.Model
}

// todo KeyMAp

var reposColumns = []table.Column{
	{Title: "Name", Width: 50},
	{Title: "Id", Width: 50},
	{Title: "URL", Width: 70},
	// todo
}

func NewReposModel(projectId string) *ReposModel {
	// r := azure.GetRepos
	// todo azure
	//testId := "002d8df7-4e1d-4c4a-9dae-83f9ea62c2cb"
	var repos = azure.GetReposForProject(projectId)

	var rows []table.Row
	for _, r := range *repos {
		row := table.Row{
			*r.Name,
			r.Id.String(),
			*r.Url,
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

	m := &ReposModel{
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
			return m, tea.Quit
		}
	}

	// Delegate message handling to the table and update its state
	var cmd tea.Cmd
	updatedTable, tableCmd := m.table.Update(msg)
	m.table = updatedTable
	cmd = tea.Batch(cmd, tableCmd)

	return m, cmd
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
