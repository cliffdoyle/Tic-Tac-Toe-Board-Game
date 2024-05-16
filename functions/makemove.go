package functions

import "fmt"

// Function to make a move on the board
func MakeMove(board [][]rune, row, col int, player rune) bool {
	if board[row][col] == ' ' {
		board[row][col] = player // Set the player's symbol in the chosen cell
		return true
	} else {
		fmt.Println("This cell is already occupied.Please choose another cell.")
		return false
	}
}
