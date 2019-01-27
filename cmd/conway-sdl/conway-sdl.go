package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crmaykish/goconway/pkg/config"
	"github.com/crmaykish/goconway/pkg/conway"

	"github.com/veandco/go-sdl2/sdl"
)

const configFileName = "config.json"

var conf config.ConwayConfig
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

	r1, g1, b1 := config.ColorChannels(conf.Colors.LivingCells.StartColor)
	r2, g2, b2 := config.ColorChannels(conf.Colors.LivingCells.EndColor)

	var red = colorChannelValue(age, r1, r2)
	var green = colorChannelValue(age, g1, g2)
	var blue = colorChannelValue(age, b1, b2)

	return red, green, blue
}

func deadCellColor(age int) (uint8, uint8, uint8) {
	r1, g1, b1 := config.ColorChannels(conf.Colors.DeadCells.StartColor)
	r2, g2, b2 := config.ColorChannels(conf.Colors.DeadCells.EndColor)

	var red = colorChannelValue(age, r1, r2)
	var green = colorChannelValue(age, g1, g2)
	var blue = colorChannelValue(age, b1, b2)

	return red, green, blue
}

func run() int {
	configFile, _ := ioutil.ReadFile(configFileName)

	conf = config.LoadConfig(configFile)

	var window *sdl.Window
	var renderer *sdl.Renderer

	var windowWidth = int32(conf.Board.Width * conf.Cells.SizeInPixels)
	var windowHeight = int32(conf.Board.Height * conf.Cells.SizeInPixels)

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
	engine := conway.CreateEngine(conf.Board.Width, conf.Board.Height)
	conway.Randomize(&engine, conf.Game.RandomFillPercent)

	// Main game loop
	for running {
		renderer.Clear()

		var background = sdl.Rect{0, 0, windowWidth, windowHeight}

		var borderR, borderG, borderB = config.ColorChannels(conf.Colors.BorderColor)

		renderer.SetDrawColor(borderR, borderG, borderB, 255)
		renderer.FillRect(&background)

		for i := 0; i < engine.BoardWidth; i++ {
			for j := 0; j < engine.BoardHeight; j++ {
				var r, g, b uint8

				if conway.CellAlive(&engine, i, j) {
					r, g, b = livingCellColor(conway.CellTimeAlive(&engine, i, j))
				} else {
					r, g, b = deadCellColor(conway.CellTimeDead(&engine, i, j))
				}

				var cellSize = conf.Cells.SizeInPixels
				var cellBorder = conf.Cells.BorderThickness

				var rect = sdl.Rect{int32((i * cellSize) + cellBorder), int32((j * cellSize) + cellBorder), int32(cellSize - (2 * cellBorder)), int32(cellSize - (2 * cellBorder))}
				renderer.SetDrawColor(r, g, b, 255)
				renderer.FillRect(&rect)
			}
		}

		// Render board and wait
		renderer.Present()
		sdl.Delay(uint32(1000 / conf.Game.Speed))

		if engine.Step == conf.Game.StepLimit {
			sdl.Delay(1000)
			conway.Reset(&engine)
			conway.Randomize(&engine, conf.Game.RandomFillPercent)
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
