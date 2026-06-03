package model
import "github.com/l4utan/ssh-portofolio/constants"

type AppState int

const (
	StateSplash AppState = iota
	StateMenu
	StateAbout
	StateProjects
	StateContact
)

type TickMsg struct{}

type Model struct {
	State    AppState
	Width    int
	Height   int
	Stars    []Star
	Tick     int
	Cursor   int
	Sections []string
}

func InitialModel(w, h int) Model {
	stars := make([]Star, constants.NumStars)
	for i := range stars {
		stars[i] = NewStar(w, h)
	}
	return Model{
		State:  StateSplash,
		Width:  w,
		Height: h,
		Stars:  stars,
		Sections: []string{
			"About",
			"Projects",
			"Contact",
		},
	}
}
