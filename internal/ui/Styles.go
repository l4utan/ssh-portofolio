package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Splash
	starColors = []lipgloss.Color{
		"#1E293B",
		"#334155",
		"#64748B",
		"#94A3B8",
	}
	nameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E2E8F0"))
	glitchStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#22D3EE"))
	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B"))
	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#334155"))

	// Menu
	menuTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E2E8F0")).
			MarginBottom(1)
	menuSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#64748B"))
	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#22D3EE"))
	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8"))
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#334155")).
			MarginTop(2)
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1E293B"))
)