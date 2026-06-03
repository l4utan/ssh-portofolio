package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	wishtea "github.com/charmbracelet/wish/bubbletea"
)

func main() {
	server, _ := wish.NewServer(
		wish.WithAddress("0.0.0.0:2222"),
		wish.WithHostKeyPath(".ssh/host_ed25519"),
		wish.WithMiddleware(wishtea.Middleware(teaHandler)),
	)
	_ = server.ListenAndServe()
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	w, h := pty.Window.Width, pty.Window.Height
	if w == 0 {
		w = 120
	}
	if h == 0 {
		h = 40
	}
	return initialModel(w, h), []tea.ProgramOption{
		tea.WithAltScreen(),
	}
}