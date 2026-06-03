package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/l4utan/ssh-portofolio/constants"
)

func TickCmd() tea.Cmd {
	return tea.Tick(constants.TickRate, func(time.Time) tea.Msg { return TickMsg{} })
}

func (m Model) Init() tea.Cmd {
	return TickCmd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		if len(m.Stars) == 0 {
			m.Stars = make([]Star, constants.NumStars)
			for i := range m.Stars {
				m.Stars[i] = NewStar(m.Width, m.Height)
			}
		}
		return m, nil

	case TickMsg:
		m.Tick++
		for i := range m.Stars {
			m.Stars[i].Update(m.Width, m.Height)
		}
		return m, TickCmd()

	case tea.KeyMsg:
		switch m.State {

		case StateSplash:
			switch msg.String() {
			case "enter", " ":
				m.State = StateMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case StateMenu:
			switch msg.String() {
			case "up", "k":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down", "j":
				if m.Cursor < len(m.Sections)-1 {
					m.Cursor++
				}
			case "enter":
				switch m.Cursor {
				case 0:
					m.State = StateAbout
				case 1:
					m.State = StateProjects
				case 2:
					m.State = StateContact
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case StateAbout, StateProjects, StateContact:
			switch msg.String() {
			case "esc", "b":
				m.State = StateMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}