package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ── Edit konten di sini ──────────────────────────────────────────────────────

type project struct {
	name  string
	desc  string
	tech  []string
	url   string
	year  string
}

var projectsList = []project{
	{
		name: "ssh-portfolio",
		desc: "This terminal — an interactive SSH portfolio built with Wish + Bubbletea.",
		tech: []string{"Go", "Wish", "Bubbletea", "Lipgloss"},
		url:  "github.com/joesluis/ssh-portfolio",
		year: "2025",
	},
	{
		name: "project-two",
		desc: "Short description of what this project does and why it matters.",
		tech: []string{"Go", "PostgreSQL"},
		url:  "github.com/joesluis/project-two",
		year: "2024",
	},
	{
		name: "project-three",
		desc: "Short description of what this project does and why it matters.",
		tech: []string{"TypeScript", "React"},
		url:  "github.com/joesluis/project-three",
		year: "2024",
	},
}

// ── View (jangan diubah kecuali mau ubah layout) ─────────────────────────────

func (m model) projectsView() string {
	w, h := m.width, m.height

	projectNameStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#22D3EE"))
	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#94A3B8"))
	techStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#64748B"))
	urlStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#334155")).
		Italic(true)
	yearStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#475569"))

	divider := dividerStyle.Render(strings.Repeat("─", 40))
	thinDiv := dividerStyle.Render(strings.Repeat("·", 40))

	var lines []string
	lines = append(lines, menuTitleStyle.Render("Projects"))
	lines = append(lines, divider)
	lines = append(lines, "")

	for i, p := range projectsList {
		header := fmt.Sprintf("%s  %s", projectNameStyle.Render(p.name), yearStyle.Render(p.year))
		lines = append(lines, header)
		lines = append(lines, descStyle.Render(p.desc))

		techRow := techStyle.Render("stack  ")
		for _, t := range p.tech {
			techRow += techStyle.Render(t+" ")
		}
		lines = append(lines, techRow)
		lines = append(lines, urlStyle.Render("↗  "+p.url))

		if i < len(projectsList)-1 {
			lines = append(lines, "")
			lines = append(lines, thinDiv)
			lines = append(lines, "")
		}
	}

	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, footerStyle.Render("esc / b → back to menu"))

	return centerBlock(lines, w, h)
}