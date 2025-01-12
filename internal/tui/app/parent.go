package app

import (
	"fmt"
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
	projectView    components.ProjectsModel
	repositoryView components.ReposModel
}

func NewParentModel() ParentModel {
	return ParentModel{
		currentView: projectView,
		// projectDetail will be initialized once we pick a project
		//projectView: components.NewProjectsTable([]string{"Proj A", "Proj B", "Proj C"}),
		projectView:    components.NewProjectsModel(),
		repositoryView: components.NewReposModel(),
	}
}

func (m ParentModel) Init() tea.Cmd {
	return nil
}

func (m ParentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case projectView:
		return m.updateProjectView(msg)
	case repositoryView:
		return m.updateRepoView(msg)
	default:
		panic("unhandled default case")
	}
	return m, nil
}

// updateProjectsTable is a helper to update our table child.
func (m ParentModel) updateProjectView(msg tea.Msg) (tea.Model, tea.Cmd) {
	// We update the projectsTable child
	newModel, cmd := m.projectView.Update(msg)
	m.projectView = newModel.(components.ProjectsModel)

	// Here we can detect if a user pressed “enter” on a project
	// and want to switch to the detail view
	// Pseudocode:
	// if user pressed enter:
	//     selectedProject := m.projectsTable.projects[m.projectsTable.selectedIndex]
	//     m.projectDetail = components.NewProjectDetail(selectedProject)
	//     m.currentView = ProjectDetailView

	// Check if the table child emitted a custom message
	switch msg := msg.(type) {
	case components.SelectProjectMsg:
		// The user selected a project, so switch to ProjectDetailView
		// Initialize projectDetail with the chosen project
		// fmt.Println("we are here")
		fmt.Println(msg.ProjectName)

		m.repositoryView = components.NewReposModel()
		m.currentView = repositoryView
		// We don’t have any new command, so return nil
		return m, nil
	}

	return m, cmd
}

func (m ParentModel) updateRepoView(msg tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := m.repositoryView.Update(msg)
	m.repositoryView = newModel.(components.ReposModel)
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
