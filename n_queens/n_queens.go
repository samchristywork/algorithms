package main

import (
	"fmt"
)

type Board struct {
	Width  int
	Height int
	Board  [][]int
}

func (board *Board)init_board(size int) {
	board.Width = size
	board.Height = size

	board.Board = make([][]int, size)

	for i := 0; i < size; i++ {
		board.Board[i] = make([]int, size)
	}
}

func main() {
	size := 8

	board := Board{}

	board.init_board(size)
}
