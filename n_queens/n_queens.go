package main

import (
	"fmt"
)

type Board struct {
	Width  int
	Height int
	Board  [][]int
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
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

func square_is_populated(board Board, row int, col int) bool {
	if row >= board.Height || col >= board.Width {
		return false
	}

	if row < 0 || col < 0 {
		return false
	}

	return board.Board[row][col] == 1
}

func row_is_safe(board Board, row int) bool {
	for i := 0; i < board.Width; i++ {
		if square_is_populated(board, row, i) {
			return false
		}
	}

	return true
}

func col_is_safe(board Board, col int) bool {
	for i := 0; i < board.Height; i++ {
		if square_is_populated(board, i, col) {
			return false
		}
	}

	return true
}

func diag_is_safe(board Board, row int, col int) bool {
	k := max(board.Width, board.Height)

	for i := 0; i < k; i++ {
		if square_is_populated(board, row-i, col-i) {
			return false
		}
		if square_is_populated(board, row+i, col+i) {
			return false
		}
		if square_is_populated(board, row-i, col+i) {
			return false
		}
		if square_is_populated(board, row+i, col-i) {
			return false
		}
	}

	return true
}

func square_is_safe(board Board, row int, col int) bool {
	return row_is_safe(board, row) &&
		col_is_safe(board, col) &&
		diag_is_safe(board, row, col)
}

func _n_queens(board Board, col int) (bool, int) {
	if col >= board.Width {
		return true, 0
	}

	for i := 0; i < board.Height; i++ {
		if square_is_safe(board, i, col) {
			board.Board[i][col] = 1

			foo, num_queens := _n_queens(board, col+1)
			if foo {
				return true, num_queens + 1
			}

			board.Board[i][col] = 0
		}
	}

	return false, 0
}

func n_queens(board Board) int {
	_, num_queens := _n_queens(board, 0)

	return num_queens
}

func main() {
	size := 8

	board := Board{}

	board.init_board(size)

	board.display()

	num_queens := n_queens(board)

	fmt.Printf("Number of queens: %d\n", num_queens)
	fmt.Println("")

	board.display()
}
