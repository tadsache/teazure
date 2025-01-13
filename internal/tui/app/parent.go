package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/tui/components"
)

type ViewState int

const (
	projectView ViewState = iota
	repositoryView
	pipelineView
	// pipelineView
	// ...
)

type ParentModel struct {
	currentView ViewState

	// child models here
	projectView    components.ProjectsModel
	repositoryView *components.ReposModel
	pipelineView   *components.PipelinesModel
}

func NewParentModel() ParentModel {
	return ParentModel{
		currentView: projectView,
		// projectDetail will be initialized once we pick a project
		projectView:    components.NewProjectsModel(),
		repositoryView: nil, // dont needto be initialized at startup
		pipelineView:   nil,
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
	case pipelineView:
		return m.updatePipelineView(msg)
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

	switch msg := msg.(type) {
	case components.SelectProjectMsg:
		// The user selected a project, so switch to the Repo View
		// Initialize repos with the chosen project
		m.repositoryView = components.NewReposModel(msg.ProjectId)
		m.currentView = repositoryView

		// We don’t have any new command, so return nil
		return m, nil
	}

	return m, cmd
}

func (m ParentModel) updateRepoView(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Perform the update on the repositoryView model
	newModel, cmd := m.repositoryView.Update(msg)

	// Assert the type and take the address
	updatedModel := newModel.(components.ReposModel)
	m.repositoryView = &updatedModel
	return m, cmd
}

func (m ParentModel) updatePipelineView(msg tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := m.pipelineView.Update(msg)

	// Assert the type and take the address
	updatedModel := newModel.(components.PipelinesModel)
	m.pipelineView = &updatedModel
	return m, cmd
}

func (m ParentModel) View() string {
	switch m.currentView {
	case projectView:
		return m.projectView.View()
	case repositoryView:
		return m.repositoryView.View()
	case pipelineView:
		return m.pipelineView.View()
	default:
		return "Unknown View"
	}
}
