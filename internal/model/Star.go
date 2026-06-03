package model

import "math/rand"

type Star struct {
	X, Y   float64
	Vx, Vy float64
	Char   rune
	Bright int
}

func NewStar(w, h int) Star {
	cx, cy := float64(w)/2, float64(h)/2
	dist := rand.Float64()*float64(h/3) + 1
	chars := []rune{'·', '•', '✦', '✧', '+', '×', '*'}
	return Star{
		X:      cx + dist*0.5,
		Y:      cy + dist*0.25,
		Vx:     (rand.Float64()*2 - 1) * 0.4,
		Vy:     (rand.Float64()*2 - 1) * 0.2,
		Char:   chars[rand.Intn(len(chars))],
		Bright: rand.Intn(4),
	}
}

func (s *Star) Update(w, h int) {
	s.X += s.Vx
	s.Y += s.Vy
	if s.X < 0 {
		s.X = float64(w)
	} else if s.X > float64(w) {
		s.X = 0
	}
	if s.Y < 0 {
		s.Y = float64(h)
	} else if s.Y > float64(h) {
		s.Y = 0
	}
}