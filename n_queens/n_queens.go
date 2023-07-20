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

func (board *Board)display() {
	fmt.Print("┌")
	for i := 0; i < board.Width*2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	for i := 0; i < board.Height; i++ {
		fmt.Print("│")
		for j := 0; j < board.Width; j++ {
			switch board.Board[i][j] {
			case 1:
				fmt.Print("Q ")
			case 0:
				fmt.Print(". ")
			default:
				fmt.Print("? ")
			}
		}
		fmt.Println("│")
	}

	fmt.Print("└")
	for i := 0; i < board.Width*2; i++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

func main() {
	size := 8

	board := Board{}

	board.init_board(size)
}
