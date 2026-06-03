package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Splash
	StarColors = []lipgloss.Color{
		"#1E293B",
		"#334155",
		"#64748B",
		"#94A3B8",
	}
	NameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E2E8F0"))
	GlitchStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#22D3EE"))
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64748B"))
	HintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#334155"))

	// Menu
	MenuTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E2E8F0")).
			MarginBottom(1)
	MenuSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#64748B"))
	SelectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#22D3EE"))
	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8"))
	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#334155")).
			MarginTop(2)
	DividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1E293B"))
)