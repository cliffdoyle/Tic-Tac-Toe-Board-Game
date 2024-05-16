package main

import (
	"fmt"

	"Tic-Tac-Toe-Board-Game/functions"
)

func main() {
	board := functions.InitializeBoard()
	// Print the initial board state
	functions.PrintBoard(board)

	for {
		// Player 1's turn
		fmt.Println("Player 1's turn (X):")
		var row, col int
		for {
			// Prompt player 1 for row and column
			fmt.Print("Enter row (0-2): ")
			fmt.Scanln(&row)
			fmt.Print("Enter column (0-2): ")
			fmt.Scanln(&col)
			if row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid input! Row and column must be between 0 and 2.Please try again")
			} else if functions.MakeMove(board, row, col, 'X') {
				// If the move is valid, break out of the inner loop
				break
			}
		}
		// Print the updated board
		functions.PrintBoard(board)
		// Check if player 1 wins
		if functions.CheckWin(board, 'X') {
			fmt.Println("Player 1 wins!")
			break
		}
		// Check for a draw
		if functions.CheckDraw(board) {
			fmt.Println("It's a draw!")
			break
		}
		// Player 2's turn
		fmt.Println("Player 2's turn (O):")
		for {
			// Prompt player 2 for row and column
			fmt.Print("Enter row (0-2): ")
			fmt.Scanln(&row)
			fmt.Print("Enter column (0-2): ")
			fmt.Scanln(&col)
			// Validate the move
			if row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid input! Row and column must be between 0 and 2. Please try again.")
			} else if functions.MakeMove(board, row, col, 'O') {
				// If the move is valid, break out of the inner loop
				break
			}
		}
		// Print the updated board
		functions.PrintBoard(board)
		// Check if player 2 wins
		if functions.CheckWin(board, 'O') {
			fmt.Println("Player 2 wins!")
			break
		}
		// Check for a draw
		if functions.CheckDraw(board) {
			fmt.Println("It's a draw!")
			break

		}
	}
}
