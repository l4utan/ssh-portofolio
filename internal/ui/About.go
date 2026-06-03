package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/l4utan/ssh-portofolio/internal/model"
)

var aboutContent = struct {
	name     string
	role     string
	bio      []string
	skills   []string
	location string
	open     bool
}{
	name: "l4utan",
	role: "Programmer",
	bio: []string{
		"I build things for the terminal and the web.",
		"Passionate about developer tooling, open source,",
		"and clean, minimal interfaces.",
	},
	skills:   []string{"Go", "Linux", "Bash", "Docker", "Kubernetes", "PostgreSQL"},
	location: "Malang, ID",
	open:     true,
}

func AboutView(m model.Model) string {
	w, h := m.Width, m.Height
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

	divider := DividerStyle.Render(strings.Repeat("─", 40))

	var lines []string
	lines = append(lines, MenuTitleStyle.Render("About"))
	lines = append(lines, divider)
	lines = append(lines, "")
	lines = append(lines, sectionTitleStyle.Render("// bio"))
	for _, l := range ac.bio {
		lines = append(lines, bioStyle.Render(l))
	}
	lines = append(lines, "")
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
	lines = append(lines, sectionTitleStyle.Render("// skills"))
	skillRow := ""
	for _, s := range ac.skills {
		skillRow += skillStyle.Render(s) + "  "
	}
	lines = append(lines, skillRow)
	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, FooterStyle.Render("esc / b → back to menu"))

	return CenterBlock(lines, w, h)
}

func CenterBlock(lines []string, w, h int) string {
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
