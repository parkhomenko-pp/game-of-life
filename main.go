package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

const (
	address = 0x3C
	width   = 128
	height  = 64

	alivePercent = 0.2 // 20% chance of being alive
)

type Board [][]bool

func NewBoard() Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return board
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

func (b Board) Draw(display ssd1306.Device) {
	for y := range b {
		for x := range b[y] {
			if b[y][x] {
				display.SetPixel(int16(x), int16(y), color.RGBA{0, 0, 0, 0})
			} else {
				display.SetPixel(int16(x), int16(y), color.RGBA{0, 0, 0, 255})
			}
		}
	}
	display.Display()
}

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: address,
		Width:   width,
		Height:  height,
	})

	display.ClearDisplay()

	for {
		board := NewBoard()
		// Randomly initialize the board
		for y := range board {
			for x := range board[y] {
				board[y][x] = rand.Float64() < alivePercent
			}
		}

		for !board.isEmpty() {
			board.Draw(display)
			board = board.Next()
			time.Sleep(time.Second)

			display.ClearDisplay()
		}
	}
}
