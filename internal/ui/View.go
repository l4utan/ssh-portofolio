package ui

import "github.com/l4utan/ssh-portofolio/internal/model"

func View(m model.Model) string {
	if m.Width == 0 || m.Height == 0 {
		return ""
	}
	switch m.State {
	case model.StateSplash:
		return SplashView(m)
	case model.StateMenu:
		return MenuView(m)
	case model.StateAbout:
		return AboutView(m)
	case model.StateProjects:
		return ProjectsView(m)
	case model.StateContact:
		return ContactView(m)
	}
	return ""
}