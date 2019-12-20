package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	windowBounds = pixel.ZV
)

func renderGame(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, imd *imdraw.IMDraw, dt float64) {
	player.render(win, viewCanvas, dt)
}

func updateGame(win *pixelgl.Window, viewCanvas *pixelgl.Canvas, dt float64) {
	player.update(win, dt)

	// why did i add this in my other engine
	if win.Bounds().W() > windowBounds.X {
		windowBounds.X = win.Bounds().W()
	}
	if win.Bounds().H() > windowBounds.Y {
		windowBounds.Y = win.Bounds().H()
	}
}
