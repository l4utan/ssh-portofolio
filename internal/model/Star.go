package model

import "math/rand"

type star struct {
	x, y   float64
	vx, vy float64
	char   rune
	bright int // 0–3
}

func newStar(w, h int) star {
	cx, cy := float64(w)/2, float64(h)/2
	angle := rand.Float64() * 2 * 3.14159
	_ = angle
	dist := rand.Float64()*float64(h/3) + 1
	chars := []rune{'·', '•', '✦', '✧', '+', '×', '*'}
	return star{
		x:      cx + dist*0.5,
		y:      cy + dist*0.25,
		vx:     (rand.Float64()*2 - 1) * 0.4,
		vy:     (rand.Float64()*2 - 1) * 0.2,
		char:   chars[rand.Intn(len(chars))],
		bright: rand.Intn(4),
	}
}

func (s *star) update(w, h int) {
	s.x += s.vx
	s.y += s.vy
	if s.x < 0 {
		s.x = float64(w)
	} else if s.x > float64(w) {
		s.x = 0
	}
	if s.y < 0 {
		s.y = float64(h)
	} else if s.y > float64(h) {
		s.y = 0
	}
}