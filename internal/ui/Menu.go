package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) menuView() string {
	w, h := m.width, m.height

	divider := dividerStyle.Render(strings.Repeat("─", 30))

	var lines []string
	lines = append(lines, menuTitleStyle.Render(yourName))
	lines = append(lines, menuSubtitleStyle.Render("Programmer"))
	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, "")

	for i, s := range m.sections {
		if i == m.cursor {
			lines = append(lines, selectedStyle.Render("▶ "+s))
		} else {
			lines = append(lines, normalStyle.Render(s))
		}
	}

	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, footerStyle.Render("↑/↓ navigate · enter select · q quit"))

	content := strings.Join(lines, "\n")
	contentLines := strings.Split(content, "\n")

	topPad := (h - len(contentLines)) / 2
	if topPad < 0 {
		topPad = 0
	}

	var sb strings.Builder
	sb.WriteString(strings.Repeat("\n", topPad))

	for _, line := range contentLines {
		// lipgloss.Width strips ANSI codes before measuring
		visLen := lipgloss.Width(line)
		leftPad := (w - visLen) / 2
		if leftPad < 0 {
			leftPad = 0
		}
		sb.WriteString(strings.Repeat(" ", leftPad))
		sb.WriteString(line)
		sb.WriteRune('\n')
	}

	return sb.String()
}