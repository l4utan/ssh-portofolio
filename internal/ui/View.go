package ui

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}
	switch m.state {
	case stateSplash:
		return m.splashView()
	case stateMenu:
		return m.menuView()
	case stateAbout:
		return m.aboutView()
	case stateProjects:
		return m.projectsView()
	case stateContact:
		return m.contactView()
	}
	return ""
}