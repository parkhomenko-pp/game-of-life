package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 100
	height = 6
)

type Board [][]bool

func NewBoard() Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return board
}

func (b Board) Print() {
	for _, row := range b {
		for _, cell := range row {
			if cell {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (b Board) Next() Board {
	newBoard := NewBoard()
	for y := range b {
		for x := range b[y] {
			aliveNeighbors := b.aliveNeighbors(x, y)
			if b[y][x] {
				newBoard[y][x] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				newBoard[y][x] = aliveNeighbors == 3
			}
		}
	}
	return newBoard
}

func (b Board) aliveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := (x+dx+width)%width, (y+dy+height)%height
			if b[ny][nx] {
				count++
			}
		}
	}
	return count
}

func (b Board) isEmpty() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell {
				return false
			}
		}
	}
	return true
}

func main() {
	for {
		board := NewBoard()
		// Randomly initialize the board
		for y := range board {
			for x := range board[y] {
				board[y][x] = rand.Float64() < 0.2 // 20% chance of being alive
			}
		}

		for !board.isEmpty() {
			board.Print()
			board = board.Next()
			time.Sleep(time.Second)
			fmt.Print("\033[H\033[2J") // Clear the screen
		}
	}
}
