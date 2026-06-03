package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(tickRate, func(time.Time) tea.Msg { return tickMsg{} })
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if len(m.stars) == 0 {
			m.stars = make([]star, numStars)
			for i := range m.stars {
				m.stars[i] = newStar(m.width, m.height)
			}
		}
		return m, nil

	case tickMsg:
		m.tick++
		for i := range m.stars {
			m.stars[i].update(m.width, m.height)
		}
		return m, tickCmd()

	case tea.KeyMsg:
		switch m.state {

		case stateSplash:
			switch msg.String() {
			case "enter", " ":
				m.state = stateMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case stateMenu:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.sections)-1 {
					m.cursor++
				}
			case "enter":
				switch m.cursor {
				case 0:
					m.state = stateAbout
				case 1:
					m.state = stateProjects
				case 2:
					m.state = stateContact
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case stateAbout, stateProjects, stateContact:
			switch msg.String() {
			case "esc", "b":
				m.state = stateMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}