package config

import (
	"encoding/json"
	"strconv"
)

type Board struct {
	Width  int
	Height int
}

type Cells struct {
	SizeInPixels    int
	BorderThickness int
}

type Game struct {
	Speed             int
	RandomFillPercent int
	StepLimit         int
}

type CellColors struct {
	StartColor string
	EndColor   string
}

type Colors struct {
	BorderColor string
	DeadCells   CellColors
	LivingCells CellColors
}

type ConwayConfig struct {
	Board  Board
	Cells  Cells
	Game   Game
	Colors Colors
}

func LoadConfig(jsonFile []byte) ConwayConfig {
	var c ConwayConfig
	json.Unmarshal(jsonFile, &c)
	return c
}

// TODO: These should be parsed once not constantly
func ColorChannels(colorString string) (uint8, uint8, uint8) {
	c, _ := strconv.ParseUint(colorString, 16, 24)

	red := (c & 0xFF0000) >> 16
	green := (c & 0xFF00) >> 8
	blue := (c & 0xFF)

	return uint8(red), uint8(green), uint8(blue)
}
