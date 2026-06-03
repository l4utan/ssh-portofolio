package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ── Edit konten di sini ──────────────────────────────────────────────────────

var aboutContent = struct {
	name     string
	role     string
	bio      []string
	skills   []string
	location string
	open     bool
}{
	name: "Joe Sluis",
	role: "Software Engineer",
	bio: []string{
		"I build things for the terminal and the web.",
		"Passionate about developer tooling, open source,",
		"and clean, minimal interfaces.",
	},
	skills:   []string{"Go", "Linux", "Bash", "Docker", "Kubernetes", "PostgreSQL"},
	location: "Amsterdam, NL",
	open:     true, // open to work?
}

// ── View (jangan diubah kecuali mau ubah layout) ─────────────────────────────

func (m model) aboutView() string {
	w, h := m.width, m.height
	ac := aboutContent

	sectionTitleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#22D3EE"))
	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#64748B"))
	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E2E8F0"))
	bioStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#94A3B8"))
	skillStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#22D3EE")).
		PaddingLeft(1).PaddingRight(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#1E293B"))
	openStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4ADE80")).Bold(true)
	closedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F87171")).Bold(true)

	divider := dividerStyle.Render(strings.Repeat("─", 40))

	var lines []string
	lines = append(lines, menuTitleStyle.Render("About"))
	lines = append(lines, divider)
	lines = append(lines, "")

	// Bio
	lines = append(lines, sectionTitleStyle.Render("// bio"))
	for _, l := range ac.bio {
		lines = append(lines, bioStyle.Render(l))
	}
	lines = append(lines, "")

	// Info
	lines = append(lines, sectionTitleStyle.Render("// info"))
	lines = append(lines, labelStyle.Render("name     ")+"  "+valueStyle.Render(ac.name))
	lines = append(lines, labelStyle.Render("role     ")+"  "+valueStyle.Render(ac.role))
	lines = append(lines, labelStyle.Render("location ")+"  "+valueStyle.Render(ac.location))
	openVal := closedStyle.Render("not available")
	if ac.open {
		openVal = openStyle.Render("open to work ✓")
	}
	lines = append(lines, labelStyle.Render("status   ")+"  "+openVal)
	lines = append(lines, "")

	// Skills
	lines = append(lines, sectionTitleStyle.Render("// skills"))
	skillRow := ""
	for _, s := range ac.skills {
		skillRow += skillStyle.Render(s) + "  "
	}
	lines = append(lines, skillRow)
	lines = append(lines, "")

	lines = append(lines, divider)
	lines = append(lines, footerStyle.Render("esc / b → back to menu"))

	return centerBlock(lines, w, h)
}

// ── Helper: center block vertically + horizontally ───────────────────────────

func centerBlock(lines []string, w, h int) string {
	content := strings.Join(lines, "\n")
	contentLines := strings.Split(content, "\n")

	topPad := (h - len(contentLines)) / 2
	if topPad < 0 {
		topPad = 0
	}

	var sb strings.Builder
	sb.WriteString(strings.Repeat("\n", topPad))
	for _, line := range contentLines {
		visLen := lipgloss.Width(line)
		leftPad := (w - visLen) / 2
		if leftPad < 0 {
			leftPad = 0
		}
		sb.WriteString(fmt.Sprintf("%s%s\n", strings.Repeat(" ", leftPad), line))
	}
	return sb.String()
}