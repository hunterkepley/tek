package main

import (
	"github.com/faiface/pixel/pixelgl"
)

//Menu ... The main menu of the game and all the UI and functions that belong to it
type Menu struct {
}

func createMainMenu() Menu {
	menu := Menu{}
	//go menu.runMusic() // Plays music
	return menu
}

func (m *Menu) update(win *pixelgl.Window, viewCanvas *pixelgl.Canvas) {
}

func (m *Menu) render(viewCanvas *pixelgl.Canvas) {
}

func (m *Menu) runMusic() {
}
