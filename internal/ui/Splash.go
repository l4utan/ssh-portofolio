package ui

import (
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) splashView() string {
	w, h := m.width, m.height

	type cell struct {
		ch    rune
		color lipgloss.Color
	}
	grid := make([][]cell, h)
	for i := range grid {
		grid[i] = make([]cell, w)
		for j := range grid[i] {
			grid[i][j] = cell{' ', ""}
		}
	}

	for _, s := range m.stars {
		x, y := int(s.x), int(s.y)
		if x >= 0 && x < w && y >= 0 && y < h {
			grid[y][x] = cell{s.char, starColors[s.bright]}
		}
	}

	nameRunes := []rune(yourName)
	nameLen := len(nameRunes)
	nameRow := h / 2
	nameCol := (w - nameLen) / 2
	if nameCol < 0 {
		nameCol = 0
	}

	subtitleRow := nameRow + 2
	hintRow := nameRow + 4

	var sb strings.Builder

	for row := 0; row < h; row++ {
		if row > 0 {
			sb.WriteRune('\n')
		}

		switch row {
		case nameRow:
			sb.WriteString(strings.Repeat(" ", nameCol))
			sb.WriteString(glitchName(nameRunes, m.tick))
			remaining := w - nameCol - nameLen
			if remaining > 0 {
				sb.WriteString(strings.Repeat(" ", remaining))
			}

		case subtitleRow:
			subtitle := "Programmer"
			col := (w - len(subtitle)) / 2
			if col < 0 {
				col = 0
			}
			sb.WriteString(strings.Repeat(" ", col))
			sb.WriteString(subtitleStyle.Render(subtitle))

		case hintRow:
			hint1 := "enter to continue"
			col1 := (w - len(hint1)) / 2
			if col1 < 0 {
				col1 = 0
			}
			sb.WriteString(strings.Repeat(" ", col1))
			sb.WriteString(hintStyle.Render(hint1))
			sb.WriteRune('\n')
			hint2 := "q / esc to quit"
			col2 := (w - len(hint2)) / 2
			if col2 < 0 {
				col2 = 0
			}
			sb.WriteString(strings.Repeat(" ", col2))
			sb.WriteString(hintStyle.Render(hint2))
			row++

		default:
			for col := 0; col < w; col++ {
				c := grid[row][col]
				if c.ch != ' ' && c.color != "" {
					sb.WriteString(lipgloss.NewStyle().Foreground(c.color).Render(string(c.ch)))
				} else {
					sb.WriteRune(' ')
				}
			}
		}
	}

	return sb.String()
}

func glitchName(runes []rune, tick int) string {
	glitchRunes := []rune(glitchChars)
	result := make([]rune, len(runes))
	copy(result, runes)

	if tick%3 == 0 {
		count := rand.Intn(3) + 1
		for i := 0; i < count; i++ {
			idx := rand.Intn(len(runes))
			if runes[idx] == ' ' {
				continue
			}
			result[idx] = glitchRunes[rand.Intn(len(glitchRunes))]
		}
	}

	var sb strings.Builder
	for i, r := range result {
		if r != runes[i] {
			sb.WriteString(glitchStyle.Render(string(r)))
		} else {
			sb.WriteString(nameStyle.Render(string(r)))
		}
	}
	return sb.String()
}