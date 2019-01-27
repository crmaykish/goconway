package main

import (
	"fmt"
	"os"

	"github.com/crmaykish/goconway/pkg/conway"

	"github.com/veandco/go-sdl2/sdl"
)

const boardWidth = 64
const boardHeight = 48
const cellPixels = 8
const cellBorder = 1
const windowWidth = boardWidth * cellPixels
const windowHeight = boardHeight * cellPixels
const speed = 8
const fill = 20
const stepLimit = 512

// Background color (really the border colors)
const bgr, bgg, bgb uint8 = 0x0F, 0x0F, 0x0F

// Dead cell colors
const dStartR, dStartG, dStartB uint8 = 0x0F, 0x41, 0x4F
const dEndR, dEndG, dEndB uint8 = 0x02, 0x2A, 0x35

// Live cell colors
const lStartR, lStartG, lStartB uint8 = 0xCA, 0xEA, 0x9C
const lEndR, lEndG, lEndB uint8 = 0x4D, 0x75, 0x14

var running = true

func colorChannelValue(age int, start, end uint8) uint8 {
	var colorChannelValue uint8

	if start < end {
		// If channel increases with age
		if age >= (int(end) - int(start)) {
			colorChannelValue = end
		} else {
			colorChannelValue = start + uint8(age) // losing data by converting int to uint8, but the range checks should protect it
		}
	} else {
		// If channel decreases with age
		if age >= (int(start) - int(end)) {
			colorChannelValue = end
		} else {
			colorChannelValue = start - uint8(age)
		}
	}

	return colorChannelValue
}

func livingCellColor(age int) (uint8, uint8, uint8) {
	// TODO: smooth the transition from all colors to take equal time

	var red = colorChannelValue(age, lStartR, lEndR)
	var green = colorChannelValue(age, lStartG, lEndG)
	var blue = colorChannelValue(age, lStartB, lEndB)

	return red, green, blue
}

func deadCellColor(age int) (uint8, uint8, uint8) {
	var red = colorChannelValue(age, dStartR, dEndR)
	var green = colorChannelValue(age, dStartG, dEndG)
	var blue = colorChannelValue(age, dStartB, dEndB)

	return red, green, blue
}

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
		renderer.SetDrawColor(bgr, bgg, bgb, 255)
		renderer.FillRect(&background)

		for i := 0; i < engine.BoardWidth; i++ {
			for j := 0; j < boardHeight; j++ {
				var r, g, b uint8

				if conway.CellAlive(&engine, i, j) {
					r, g, b = livingCellColor(conway.CellTimeAlive(&engine, i, j))
				} else {
					r, g, b = deadCellColor(conway.CellTimeDead(&engine, i, j))
				}

				var rect = sdl.Rect{int32(i*cellPixels) + cellBorder, int32(j*cellPixels) + cellBorder, cellPixels - (2 * cellBorder), cellPixels - (2 * cellBorder)}
				renderer.SetDrawColor(r, g, b, 255)
				renderer.FillRect(&rect)
			}
		}

		// Render board and wait
		renderer.Present()
		sdl.Delay(1000 / speed)

		if engine.Step == stepLimit {
			sdl.Delay(1000)
			conway.Reset(&engine)
			conway.Randomize(&engine, fill)
		} else {
			// Process the next step in the game
			conway.Step(&engine)
		}

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

func main() {
	os.Exit(run())
}
