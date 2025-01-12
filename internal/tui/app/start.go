package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// todo add cfg app.Config as arg
func Start() error {
	//m := projectsModel()
	m := NewParentModel()
	a := tea.NewProgram(m,
		tea.WithAltScreen())
	// add routines here
	_, err := a.Run()
	return err
}
