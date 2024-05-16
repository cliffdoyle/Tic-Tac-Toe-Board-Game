package functions

// Function to check if the game is a draw
func CheckDraw(board [][]rune) bool {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == ' ' {
				return false // if any cell is empty, the game is not a draw yet
			}
		}
	}
	return true // If all cells are occupied and there's no winner, the game is a draw
}
