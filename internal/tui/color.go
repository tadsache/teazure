package tui

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// todo config for the theme ..
//
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

var GlobalTheme Theme

var (
// safe vars that are used in the tui here
)

// a Theme can be load by a ghostty theme file
// maynbe make more file types accessible
// todo think about path and config how to make themes available
// for testing purposes ghostty app paths Applications/Ghostty.app/Contents/Resources/ghostty/themes

func LoadTheme() error {
	themeName := os.Getenv("THEME_NAME") // fixme need to be in config and can be as a arg at startup
	//themeName := "catppuccin-mocha"
	themeBasePath := os.Getenv("THEME_BASE_PATH")
	themePath := fmt.Sprintf("%s/%s", themeBasePath, themeName)
	t, err := parseThemeFile(themePath)
	if err != nil {
		return err
	}
	GlobalTheme = t
	return nil // is this good practice ?
}

func parseThemeFile(path string) (Theme, error) {
	var t Theme

	f, err := os.Open(path)
	if err != nil {
		return t, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Regex to match lines like "palette = 0=#494d64"
	// Capture group 1 => index (0..15)
	// Capture group 2 => hex color (#494d64)
	paletteRegex := regexp.MustCompile(`^palette\s*=\s*(\d+)\s*=\s*(#[A-Fa-f0-9]{6})`)

	// Regex to match lines like "background = #24273a"
	// Capture group 1 => key (background, foreground, etc.)
	// Capture group 2 => hex color (#24273a)
	keyValueRegex := regexp.MustCompile(`^(background|foreground|cursor-color|selection-background|selection-foreground)\s*=\s*(#[A-Fa-f0-9]{6})`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			// skip empty lines or commented lines
			continue
		}

		// Check if it's a "palette = X=#XXXXXX" line
		if matches := paletteRegex.FindStringSubmatch(line); len(matches) == 3 {
			// e.g., matches[1] = "0", matches[2] = "#494d64"
			idxStr := matches[1]
			colorHex := matches[2]
			idx, _ := strconv.Atoi(idxStr)

			switch idx {
			case 0:
				t.Black = lipgloss.Color(colorHex)
			case 1:
				t.Red = lipgloss.Color(colorHex)
			case 2:
				t.Green = lipgloss.Color(colorHex)
			case 3:
				t.Yellow = lipgloss.Color(colorHex)
			case 4:
				t.Blue = lipgloss.Color(colorHex)
			case 5:
				t.Purple = lipgloss.Color(colorHex)
			case 6:
				t.Cyan = lipgloss.Color(colorHex)
			case 7:
				t.White = lipgloss.Color(colorHex)
			case 8:
				t.BrightBlack = lipgloss.Color(colorHex)
			case 9:
				t.BrightRed = lipgloss.Color(colorHex)
			case 10:
				t.BrightGreen = lipgloss.Color(colorHex)
			case 11:
				t.BrightYellow = lipgloss.Color(colorHex)
			case 12:
				t.BrightBlue = lipgloss.Color(colorHex)
			case 13:
				t.BrightPurple = lipgloss.Color(colorHex)
			case 14:
				t.BrightCyan = lipgloss.Color(colorHex)
			case 15:
				t.BrightWhite = lipgloss.Color(colorHex)
			}
			continue
		}

		// Otherwise, check if it's one of the special keys
		if matches := keyValueRegex.FindStringSubmatch(line); len(matches) == 3 {
			key := matches[1]
			colorHex := matches[2]

			switch key {
			case "background":
				t.Background = lipgloss.Color(colorHex)
			case "foreground":
				t.Foreground = lipgloss.Color(colorHex)
			case "cursor-color":
				t.Cursor = lipgloss.Color(colorHex)
			case "selection-background":
				t.SelectionBackground = lipgloss.Color(colorHex)
			case "selection-foreground":
				t.SelectionForeground = lipgloss.Color(colorHex)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return t, fmt.Errorf("error reading file: %w", err)
	}

	return t, nil
}
