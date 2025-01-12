package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/tui/components"
)

type ViewState int

const (
	projectView ViewState = iota
	repositoryView
	// pipelineView
	// ...
)

type ParentModel struct {
	currentView ViewState

	// child models here
	projectView components.ProjectModel
}

func NewParentModel() ParentModel {
	return ParentModel{
		currentView: projectView,
		// projectDetail will be initialized once we pick a project
		//projectView: components.NewProjectsTable([]string{"Proj A", "Proj B", "Proj C"}),
		projectView: components.NewAzProjectsTable(),
	}
}

func (m ParentModel) Init() tea.Cmd {
	return nil
}

func (m ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case projectView:
		return m.updateProjectView(msg)

	default:
		panic("unhandled default case")
	}
	return m, nil
}

// updateProjectsTable is a helper to update our table child.
func (m ParentModel) updateProjectView(msg tea.Msg) (tea.Model, tea.Cmd) {
	// We update the projectsTable child
	newTableModel, cmd := m.projectView.Update(msg)
	m.projectView = newTableModel.(components.ProjectModel)

	// Here we can detect if a user pressed “enter” on a project
	// and want to switch to the detail view
	// Pseudocode:
	// if user pressed enter:
	//     selectedProject := m.projectsTable.projects[m.projectsTable.selectedIndex]
	//     m.projectDetail = components.NewProjectDetail(selectedProject)
	//     m.currentView = ProjectDetailView

	return m, cmd
}

func (m ParentModel) View() string {
	switch m.currentView {
	case projectView:
		return m.projectView.View()
	default:
		return "Unknown View"
	}
}
