package main

import (
	"fmt"

	"os"

	"unicode"
)


func solveSudoku(board [][]int) bool {

	row, col, found := findEmpty(board)

	if !found {

		return true 

	}

	for val := 1; val <= 9; val++ {

		if isValid(board, row, col, val) {

			board[row][col] = val

			if solveSudoku(board) {

				return true

			}

			board[row][col] = 0 

		}

	}

	return false

}


func findEmpty(board [][]int) (int, int, bool) {

	for i := 0; i < 9; i++ {

		for j := 0; j < 9; j++ {

			if board[i][j] == 0 {

				return i, j, true

			}

		}

	}

	return -1, -1, false

}


func isValid(board [][]int, row, col, val int) bool {

	for i := 0; i < 9; i++ {

		if board[row][i] == val || board[i][col] == val {

			return false

		}

	}

	boxRowStart := (row / 3) * 3

	boxColStart := (col / 3) * 3

	for r := boxRowStart; r < boxRowStart+3; r++ {

		for c := boxColStart; c < boxColStart+3; c++ {

			if board[r][c] == val {

				return false

			}

		}

	}

	return true

}



func parseInput(args []string) ([][]int, error) {

	if len(args) != 9 {

		return nil, fmt.Errorf("invalid number of rows")

	}

	board := make([][]int, 9)

	for i, row := range args {

		if len(row) != 9 {

			return nil, fmt.Errorf("invalid row length")

		}

		board[i] = make([]int, 9)

		for j, ch := range row {

			if ch == '.' {

				board[i][j] = 0

			} else if unicode.IsDigit(ch) && ch != '0' {

				board[i][j] = int(ch - '0')

			} else {

				return nil, fmt.Errorf("invalid character")

			}

		}

	}

	return board, nil

}


func printBoard(board [][]int) {

	for i := 0; i < 9; i++ {

		for j := 0; j < 9; j++ {

			fmt.Printf("%d", board[i][j])

			if j < 8 {

				fmt.Print(" ")

			}

		}

		fmt.Println()

	}

}

func main() {

	args := os.Args[1:]

	board, err := parseInput(args)

	if err != nil {

		fmt.Println("Error")

		return

	}

	if !solveSudoku(board) {

		fmt.Println("Error")

		return

	}

	printBoard(board)

}
