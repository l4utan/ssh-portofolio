package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/l4utan/ssh-portofolio/constants"
	"github.com/l4utan/ssh-portofolio/internal/model"
)

func MenuView(m model.Model) string {
	w, h := m.Width, m.Height

	divider := DividerStyle.Render(strings.Repeat("─", 30))

	var lines []string
	lines = append(lines, MenuTitleStyle.Render(constants.YourName))
	lines = append(lines, MenuSubtitleStyle.Render("Programmer"))
	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, "")

	for i, s := range m.Sections {
		if i == m.Cursor {
			lines = append(lines, SelectedStyle.Render("▶ "+s))
		} else {
			lines = append(lines, NormalStyle.Render(s))
		}
	}

	lines = append(lines, "")
	lines = append(lines, divider)
	lines = append(lines, FooterStyle.Render("↑/↓ navigate · enter select · q quit"))

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
		sb.WriteString(strings.Repeat(" ", leftPad))
		sb.WriteString(line)
		sb.WriteRune('\n')
	}

	return sb.String()
}