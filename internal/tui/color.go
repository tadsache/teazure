package tui

import "github.com/charmbracelet/lipgloss"

// todo config for the theme ..
//
//	/Applications/Ghostty.app/Contents/Resources/ghostty/the
//	/Applications/Ghostty.app/Contents/Resources/ghostty/themes
type Theme struct {
	Black               lipgloss.Color // pallete 0
	Red                 lipgloss.Color // pallete 1
	Green               lipgloss.Color // pallete 2
	Yellow              lipgloss.Color // pallete 3
	Blue                lipgloss.Color // pallete 4
	Purple              lipgloss.Color // pallete 5
	Cyan                lipgloss.Color // pallete 6
	White               lipgloss.Color // pallete 7
	BrightBlack         lipgloss.Color // pallete 8
	BrightRed           lipgloss.Color // pallete 9
	BrightGreen         lipgloss.Color // pallete 10
	BrightYellow        lipgloss.Color // pallete 11
	BrightBlue          lipgloss.Color // pallete 12
	BrightPurple        lipgloss.Color // pallete 13
	BrightCyan          lipgloss.Color // pallete 14
	BrightWhite         lipgloss.Color // pallete 15
	Background          lipgloss.Color
	Foreground          lipgloss.Color
	Cursor              lipgloss.Color
	SelectionBackground lipgloss.Color
	SelectionForeground lipgloss.Color
}

var (
// safe vars that are used in the tui here
)
