package engine

import (
	"fmt"
	"math/rand"
	"time"
)

type conwayCell struct {
	CurrentlyAlive  bool
	PreviouslyAlive bool
}

type ConwayEngine struct {
	BoardWidth  int
	BoardHeight int
	cells       [][]conwayCell
}

var BoardWidth = 20
var BoardHeight = 15

func InitBoard(x, y int) [][]Cell {
	BoardWidth = x
	BoardHeight = y

	board := make([][]Cell, x)

	for i := 0; i < x; i++ {
		board[i] = make([]Cell, y)
	}

	return board
}

func PrintBoard(b [][]Cell) {
	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			if b[j][i].CurrentlyAlive {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" \n")
	}
}

// TODO: parallelize this
func Step(b [][]Cell) {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			var neighbors = livingNeighbors(b, i, j)

			if b[i][j].PreviouslyAlive { // Living Cell rules
				// Cell dies
				if neighbors < 2 || neighbors > 3 {
					b[i][j].CurrentlyAlive = false
				}

				// Cells lives
				if neighbors >= 2 && neighbors <= 3 {
					b[i][j].CurrentlyAlive = true
				}

			} else { // Dead Cell rules
				if neighbors == 3 {
					b[i][j].CurrentlyAlive = true
				} else {
					b[i][j].CurrentlyAlive = false
				}
			}
		}
	}

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			b[i][j].PreviouslyAlive = b[i][j].CurrentlyAlive
		}
	}
}

func Randomize(b [][]Cell, fillPercent int) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			if r.Intn(100) > (100 - fillPercent) {
				b[i][j].PreviouslyAlive = true
				b[i][j].CurrentlyAlive = true
			}
		}
	}
}

func livingNeighbors(board [][]Cell, x int, y int) int {
	var alive = 0

	// top left
	if x > 0 && y > 0 && board[x-1][y-1].PreviouslyAlive {
		alive++
	}

	// top
	if y > 0 && board[x][y-1].PreviouslyAlive {
		alive++
	}

	// top right
	if x < BoardWidth-1 && y > 0 && board[x+1][y-1].PreviouslyAlive {
		alive++
	}

	// left
	if x > 0 && board[x-1][y].PreviouslyAlive {
		alive++
	}

	// right
	if x < BoardWidth-1 && board[x+1][y].PreviouslyAlive {
		alive++
	}

	// bottom left
	if x > 0 && y < BoardHeight-1 && board[x-1][y+1].PreviouslyAlive {
		alive++
	}

	// bottom
	if y < BoardHeight-1 && board[x][y+1].PreviouslyAlive {
		alive++
	}

	// bottom right
	if x < BoardWidth-1 && y < BoardHeight-1 && board[x+1][y+1].PreviouslyAlive {
		alive++
	}

	return alive
}
