package main

import (
	"fmt"
	"time"

	"github.com/crmaykish/goconway/pkg/conway"
)

var boardWidth = 15
var boardHeight = 10
var speed = 2
var fillPercent = 30

func main() {
	fmt.Println("Starting Game of Life...")

	engine := conway.CreateEngine(boardWidth, boardHeight)

	conway.Randomize(&engine, fillPercent)

	conway.PrintBoard(&engine)

	for {
		conway.Step(&engine)
		conway.PrintBoard(&engine)
		time.Sleep(time.Millisecond * time.Duration(1000/speed))
	}
}
