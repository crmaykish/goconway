package conway

import (
	"fmt"
	"math/rand"
	"time"
)

type conwayCell struct {
	TimeAlive       int
	TimeDead        int
	Alive           bool
	previouslyAlive bool
}

type ConwayEngine struct {
	Step        int
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
			if e.board[j][i].Alive {
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

			if e.board[i][j].previouslyAlive { // Living Cell rules
				// Cell dies
				if neighbors < 2 || neighbors > 3 {
					e.board[i][j].Alive = false
					e.board[i][j].TimeAlive = 0
				}

				// Cells stays alive
				if neighbors >= 2 && neighbors <= 3 {
					e.board[i][j].Alive = true
					e.board[i][j].TimeAlive++
					e.board[i][j].TimeDead = 0
				}

			} else { // Dead Cell rules
				if neighbors == 3 {
					// Cell comes alive
					e.board[i][j].Alive = true
				} else {
					// Cell stays dead
					e.board[i][j].Alive = false
					e.board[i][j].TimeDead++
				}
			}
		}
	}

	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			e.board[i][j].previouslyAlive = e.board[i][j].Alive
		}
	}

	e.Step++
}

func Randomize(e *ConwayEngine, fillPercent int) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			if r.Intn(100) > (100 - fillPercent) {
				e.board[i][j].previouslyAlive = true
				e.board[i][j].Alive = true
			}
		}
	}
}

func Reset(e *ConwayEngine) {
	e.Step = 0
	for i := 0; i < e.BoardWidth; i++ {
		for j := 0; j < e.BoardHeight; j++ {
			e.board[i][j].TimeAlive = 0
			e.board[i][j].TimeDead = 0
			e.board[i][j].Alive = false
			e.board[i][j].previouslyAlive = false
		}
	}
}

func CellAlive(e *ConwayEngine, x, y int) bool {
	return e.board[x][y].Alive
}

func CellTimeAlive(e *ConwayEngine, x, y int) int {
	return e.board[x][y].TimeAlive
}

func CellTimeDead(e *ConwayEngine, x, y int) int {
	return e.board[x][y].TimeDead
}

func livingNeighbors(e *ConwayEngine, x int, y int) int {
	var alive = 0

	// top left
	if x > 0 && y > 0 && e.board[x-1][y-1].previouslyAlive {
		alive++
	}

	// top
	if y > 0 && e.board[x][y-1].previouslyAlive {
		alive++
	}

	// top right
	if x < e.BoardWidth-1 && y > 0 && e.board[x+1][y-1].previouslyAlive {
		alive++
	}

	// left
	if x > 0 && e.board[x-1][y].previouslyAlive {
		alive++
	}

	// right
	if x < e.BoardWidth-1 && e.board[x+1][y].previouslyAlive {
		alive++
	}

	// bottom left
	if x > 0 && y < e.BoardHeight-1 && e.board[x-1][y+1].previouslyAlive {
		alive++
	}

	// bottom
	if y < e.BoardHeight-1 && e.board[x][y+1].previouslyAlive {
		alive++
	}

	// bottom right
	if x < e.BoardWidth-1 && y < e.BoardHeight-1 && e.board[x+1][y+1].previouslyAlive {
		alive++
	}

	return alive
}
