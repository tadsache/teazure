package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// default styles for bubble or lipgloss components

var TableStyle table.Styles

// LoadStyles initialize all Styles
func LoadStyles() {
	TableStyle = loadTableStyle()
}

// needs to be initialized after the Theme is loaded
func loadTableStyle() table.Styles {
	return table.Styles{
		Header: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(GlobalTheme.Purple).
			BorderBottom(true).
			Bold(false),
		Selected: lipgloss.NewStyle().
			Foreground(GlobalTheme.SelectionForeground).
			Background(GlobalTheme.SelectionBackground).
			Bold(false),
	}

}
