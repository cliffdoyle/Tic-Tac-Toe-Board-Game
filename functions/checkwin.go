package functions

// Function to check if a player has won
func CheckWin(board [][]rune, player rune) bool {
	// check rows

	for i := 0; i < rows; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// Check columns

	for j := 0; j < cols; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}

	// Check diagonals

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}
