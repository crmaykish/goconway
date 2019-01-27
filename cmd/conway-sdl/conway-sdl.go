package main

import (
	"fmt"
	"os"

	"github.com/crmaykish/goconway/pkg/engine"
	"github.com/veandco/go-sdl2/sdl"
)

const boardWidth = 500
const boardHeight = 300
const cellPixels = 4
const windowWidth = boardWidth * cellPixels
const windowHeight = boardHeight * cellPixels
const speed = 8
const fill = 10

var board [][]engine.Cell

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow("Conway's Game of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	board = engine.InitBoard(boardWidth, boardHeight)
	engine.Randomize(board, fill)

	for {
		renderer.Clear()

		var background = sdl.Rect{0, 0, windowWidth, windowHeight}
		renderer.SetDrawColor(0x00, 0x40, 0x11, 255)
		renderer.FillRect(&background)

		for i := 0; i < engine.BoardWidth; i++ {
			for j := 0; j < boardHeight; j++ {
				if board[i][j].CurrentlyAlive {
					var rect = sdl.Rect{int32(i * cellPixels), int32(j * cellPixels), cellPixels, cellPixels}
					renderer.SetDrawColor(0x4F, 0x9F, 0x64, 255)
					renderer.FillRect(&rect)
				}
			}
		}

		renderer.Present()

		engine.Step(board)
		sdl.Delay(1000 / speed)
	}
}

func main() {
	os.Exit(run())
}
