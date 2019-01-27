package main

import (
	"fmt"
	"time"

	"github.com/crmaykish/goconway/pkg/engine"
)

var speed = 2
var fillPercent = 30

func main() {
	fmt.Println("Starting Game of Life...")

	var board = make([][]engine.Cell, engine.BoardWidth)

	for i := 0; i < engine.BoardWidth; i++ {
		board[i] = make([]engine.Cell, engine.BoardHeight)
	}

	engine.Randomize(board, fillPercent)

	engine.PrintBoard(board)

	for {
		engine.Step(board)
		engine.PrintBoard(board)
		time.Sleep(time.Millisecond * time.Duration(1000/speed))
	}
}
