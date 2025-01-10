package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// default styles for bubble or lipgloss components

var TableStyle = table.Styles{
	Header: lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(rosewater).
		BorderBottom(true).
		Bold(false),
	Selected: lipgloss.NewStyle().
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#7287fd")).
		Bold(false),
}
