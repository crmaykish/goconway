package main

import (
	"fmt"
	"os"

	"github.com/crmaykish/goconway/pkg/conway"

	"github.com/veandco/go-sdl2/sdl"
)

const boardWidth = 500
const boardHeight = 300
const cellPixels = 4
const windowWidth = boardWidth * cellPixels
const windowHeight = boardHeight * cellPixels
const speed = 8
const fill = 10

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	// Create the main SDL window
	window, err := sdl.CreateWindow("Conway's Game of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	// Create the SDL renderer
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	// Create the game engine
	engine := conway.CreateEngine(boardWidth, boardHeight)
	conway.Randomize(&engine, fill)

	for {
		renderer.Clear()

		var background = sdl.Rect{0, 0, windowWidth, windowHeight}
		renderer.SetDrawColor(0x00, 0x40, 0x11, 255)
		renderer.FillRect(&background)

		for i := 0; i < engine.BoardWidth; i++ {
			for j := 0; j < boardHeight; j++ {
				if conway.CellAliveAt(&engine, i, j) {
					var rect = sdl.Rect{int32(i * cellPixels), int32(j * cellPixels), cellPixels, cellPixels}
					renderer.SetDrawColor(0x4F, 0x9F, 0x64, 255)
					renderer.FillRect(&rect)
				}
			}
		}

		// Render board and wait
		renderer.Present()
		sdl.Delay(1000 / speed)

		// Process the next step in the game
		conway.Step(&engine)
	}
}

func main() {
	os.Exit(run())
}
