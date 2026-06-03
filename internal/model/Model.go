package model

type appState int

const (
	StateSplash appState = iota
	StateMenu
	StateAbout
	StateProjects
	StateContact
)

type TickMsg struct{}

type Model struct {
	state    appState
	width    int
	height   int
	stars    []star
	tick     int
	cursor   int
	sections []string
}

func initialModel(w, h int) model {
	stars := make([]star, numStars)
	for i := range stars {
		stars[i] = newStar(w, h)
	}
	return model{
		state:  stateSplash,
		width:  w,
		height: h,
		stars:  stars,
		sections: []string{
			"About",
			"Projects",
			"Contact",
		},
	}
}