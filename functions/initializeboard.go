package functions

// Function to initialize a new Tic-Tac-Toe board.
func InitializeBoard() [][]rune {
	board := make([][]rune, rows)
	for i := range board {
		board[i] = make([]rune, cols)
		for j := range board[i] {
			board[i][j] = ' ' // Empty space represents an unoccupied cell
		}
	}
	return board
}
