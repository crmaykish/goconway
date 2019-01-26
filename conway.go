package main

import (
	"fmt"
	"math/rand"
	"time"
)

const boardWidth = 20
const boardHeight = 15

var speed = 2

type cell struct {
	CurrentlyAlive  bool
	PreviouslyAlive bool
}

var board [][]cell

func printBoard(b [][]cell) {
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			if b[j][i].CurrentlyAlive {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" \n")
	}
}

func livingNeighbors(board [][]cell, x int, y int) int {
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
	if x < 9 && y > 0 && board[x+1][y-1].PreviouslyAlive {
		alive++
	}

	// left
	if x > 0 && board[x-1][y].PreviouslyAlive {
		alive++
	}

	// right
	if x < 9 && board[x+1][y].PreviouslyAlive {
		alive++
	}

	// bottom left
	if x > 0 && y < 9 && board[x-1][y+1].PreviouslyAlive {
		alive++
	}

	// bottom
	if y < 9 && board[x][y+1].PreviouslyAlive {
		alive++
	}

	// bottom right
	if x < 9 && y < 9 && board[x+1][y+1].PreviouslyAlive {
		alive++
	}

	return alive
}

func step(b [][]cell) {
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			var neighbors = livingNeighbors(b, i, j)

			if b[i][j].PreviouslyAlive { // Living cell rules
				// Cell dies
				if neighbors < 2 || neighbors > 3 {
					b[i][j].CurrentlyAlive = false
				}

				// Cells lives
				if neighbors >= 2 && neighbors <= 3 {
					b[i][j].CurrentlyAlive = true
				}

			} else { // Dead cell rules
				if neighbors == 3 {
					b[i][j].CurrentlyAlive = true
				} else {
					b[i][j].CurrentlyAlive = false
				}
			}
		}
	}

	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			b[i][j].PreviouslyAlive = b[i][j].CurrentlyAlive
		}
	}
}

func randomize(b [][]cell) {
	var source = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(source)

	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			if r.Intn(100) > 50 {
				b[i][j].PreviouslyAlive = true
				b[i][j].CurrentlyAlive = true
			}
		}
	}
}

func main() {
	fmt.Println("Starting Game of Life...")

	board := make([][]cell, boardWidth)

	for i := 0; i < boardWidth; i++ {
		board[i] = make([]cell, boardHeight)
	}

	randomize(board)

	printBoard(board)

	for {
		step(board)
		printBoard(board)
		time.Sleep(time.Millisecond * time.Duration(1000/speed))
	}

}
