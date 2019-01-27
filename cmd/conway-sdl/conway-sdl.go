package main

import (
	"fmt"
	"os"

	"github.com/crmaykish/goconway/pkg/conway"

	"github.com/veandco/go-sdl2/sdl"
)

const boardWidth = 50
const boardHeight = 30
const cellPixels = 20
const cellBorder = 2
const windowWidth = boardWidth * cellPixels
const windowHeight = boardHeight * cellPixels
const speed = 8
const fill = 25

var running = true

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

	// Main game loop
	for running {
		renderer.Clear()

		var background = sdl.Rect{0, 0, windowWidth, windowHeight}
		renderer.SetDrawColor(0x1A, 0x1A, 0x1A, 255)
		renderer.FillRect(&background)

		for i := 0; i < engine.BoardWidth; i++ {
			for j := 0; j < boardHeight; j++ {
				if conway.CellAlive(&engine, i, j) {
					var rect = sdl.Rect{int32(i*cellPixels) + cellBorder, int32(j*cellPixels) + cellBorder, cellPixels - (2 * cellBorder), cellPixels - (2 * cellBorder)}
					var r, g, b = cellColor(conway.CellAge(&engine, i, j))
					renderer.SetDrawColor(r, g, b, 255)
					renderer.FillRect(&rect)
				}
			}
		}

		// Render board and wait
		renderer.Present()
		sdl.Delay(1000 / speed)

		// Process the next step in the game
		conway.Step(&engine)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}

	return 0
}

func cellColor(age int) (uint8, uint8, uint8) {
	var red = 0
	var green = 0
	var blue = 0

	if age >= 255 {
		red = 255
		green = 0
	} else {
		red = age
		green = 255 - age
	}

	return uint8(red), uint8(green), uint8(blue)
}

func main() {
	os.Exit(run())
}
