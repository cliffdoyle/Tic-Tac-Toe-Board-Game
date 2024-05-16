package functions

import "fmt"

const (
	rows = 3
	cols = 3
)

// function to print the Tic-Tac-Toe board
func PrintBoard(board [][]rune) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%c", board[i][j]) // print each cell with its current content
		}
		fmt.Println() // move to the next row
	}
}
