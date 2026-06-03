package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ── Edit konten di sini ──────────────────────────────────────────────────────

var contactContent = struct {
	email   string
	github  string
	twitter string
	web     string
	note    string
}{
	email:   "joe@joesluis.dev",
	github:  "github.com/joesluis",
	twitter: "@joesluis",
	web:     "joesluis.dev",
	note:    "Best reached via email. I usually reply within 24 hours.",
}

// ── View (jangan diubah kecuali mau ubah layout) ─────────────────────────────

func (m model) contactView() string {
	w, h := m.width, m.height
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

	divider := dividerStyle.Render(strings.Repeat("─", 40))

	row := func(icon, label, val string) string {
		return iconStyle.Render(icon+"  "+label+"  ") + valueStyle.Render(val)
	}

	var lines []string
	lines = append(lines, menuTitleStyle.Render("Contact"))
	lines = append(lines, divider)
	lines = append(lines, "")
	lines = append(lines, sectionTitleStyle.Render("// reach me at"))
	lines = append(lines, "")
	lines = append(lines, row("✉", "email  ", cc.email))
	lines = append(lines, row("⌥", "github ", cc.github))
	lines = append(lines, row("◈", "twitter", cc.twitter))
	lines = append(lines, row("◉", "web    ", cc.web))
	lines = append(lines, "")
	lines = append(lines, noteStyle.Render(cc.note))
	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, footerStyle.Render("esc / b → back to menu"))

	return centerBlock(lines, w, h)
}