package conway

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
	board       [][]conwayCell
}

func CreateEngine(width, height int) ConwayEngine {
	e := ConwayEngine{BoardWidth: width, BoardHeight: height}

	e.board = make([][]conwayCell, width)

	for i := 0; i < width; i++ {
		e.board[i] = make([]conwayCell, height)
	}

	return e
}

func PrintBoard(e *ConwayEngine) {
	for i := 0; i < e.BoardHeight; i++ {
		for j := 0; j < e.BoardWidth; j++ {
			if e.board[j][i].CurrentlyAlive {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" \n")
	}
}

// TODO: parallelize this
func Step(e *ConwayEngine) {
	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			var neighbors = livingNeighbors(e, i, j)

			if e.board[i][j].PreviouslyAlive { // Living Cell rules
				// Cell dies
				if neighbors < 2 || neighbors > 3 {
					e.board[i][j].CurrentlyAlive = false
				}

				// Cells lives
				if neighbors >= 2 && neighbors <= 3 {
					e.board[i][j].CurrentlyAlive = true
				}

			} else { // Dead Cell rules
				if neighbors == 3 {
					e.board[i][j].CurrentlyAlive = true
				} else {
					e.board[i][j].CurrentlyAlive = false
				}
			}
		}
	}

	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			e.board[i][j].PreviouslyAlive = e.board[i][j].CurrentlyAlive
		}
	}
}

func Randomize(e *ConwayEngine, fillPercent int) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			if r.Intn(100) > (100 - fillPercent) {
				e.board[i][j].PreviouslyAlive = true
				e.board[i][j].CurrentlyAlive = true
			}
		}
	}
}

func CellAliveAt(e *ConwayEngine, x, y int) bool {
	return e.board[x][y].CurrentlyAlive
}

func livingNeighbors(e *ConwayEngine, x int, y int) int {
	var alive = 0

	// top left
	if x > 0 && y > 0 && e.board[x-1][y-1].PreviouslyAlive {
		alive++
	}

	// top
	if y > 0 && e.board[x][y-1].PreviouslyAlive {
		alive++
	}

	// top right
	if x < e.BoardWidth-1 && y > 0 && e.board[x+1][y-1].PreviouslyAlive {
		alive++
	}

	// left
	if x > 0 && e.board[x-1][y].PreviouslyAlive {
		alive++
	}

	// right
	if x < e.BoardWidth-1 && e.board[x+1][y].PreviouslyAlive {
		alive++
	}

	// bottom left
	if x > 0 && y < e.BoardHeight-1 && e.board[x-1][y+1].PreviouslyAlive {
		alive++
	}

	// bottom
	if y < e.BoardHeight-1 && e.board[x][y+1].PreviouslyAlive {
		alive++
	}

	// bottom right
	if x < e.BoardWidth-1 && y < e.BoardHeight-1 && e.board[x+1][y+1].PreviouslyAlive {
		alive++
	}

	return alive
}
