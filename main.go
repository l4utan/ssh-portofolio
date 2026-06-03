package main

import (
	"github.com/l4utan/ssh-portofolio/internal/model"
	"github.com/l4utan/ssh-portofolio/internal/server"
	"github.com/l4utan/ssh-portofolio/internal/ui"
)

func main() {
	model.ViewFunc = ui.View
	server.Start()
}