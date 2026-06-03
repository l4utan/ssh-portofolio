package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/l4utan/ssh-portofolio/internal/model"
)

var contactContent = struct {
	email   string
	github  string
	twitter string
	web     string
	note    string
}{
	email:   "joe@joesluis.dev",
	github:  "github.com/l4utan",
	web:     "-- WORKING ON --",
	note:    "Best reached via email. I usually reply within 24 hours.",
}

func ContactView(m model.Model) string {
	w, h := m.Width, m.Height
	cc := contactContent

	sectionTitleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#22D3EE"))
	iconStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#64748B"))
	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E2E8F0"))
	noteStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#64748B")).
		Italic(true)

	divider := DividerStyle.Render(strings.Repeat("─", 40))

	row := func(icon, label, val string) string {
		return iconStyle.Render(icon+"  "+label+"  ") + valueStyle.Render(val)
	}

	var lines []string
	lines = append(lines, MenuTitleStyle.Render("Contact"))
	lines = append(lines, divider)
	lines = append(lines, "")
	lines = append(lines, sectionTitleStyle.Render("// reach me at"))
	lines = append(lines, "")
	lines = append(lines, row("✉", "email  ", cc.email))
	lines = append(lines, row("⌥", "github ", cc.github))
	lines = append(lines, row("◉", "web    ", cc.web))
	lines = append(lines, "")
	lines = append(lines, noteStyle.Render(cc.note))
	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, FooterStyle.Render("esc / b → back to menu"))

	return CenterBlock(lines, w, h)
}