// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

func isValidSudoku(board [][]byte) bool {
	sb := make(map[string]bool)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == "." {
				continue
			}

			_, okRow := sb[fmt.Sprintf("%s contains in row %s", string(board[i][j]), string(i))]
			_, okCol := sb[fmt.Sprintf("%s contains in column %s", string(board[i][j]), string(j))]
			_, okGrid3x3 := sb[fmt.Sprintf("%s contains in grid 3x3 %s %s", string(board[i][j]), string(i/3), string(j/3))]

			if okRow || okCol || okGrid3x3 {
				return false
			} else {
				sb[fmt.Sprintf("%s contains in grid 3x3 %s %s", string(board[i][j]), string(i/3), string(j/3))] = true
				sb[fmt.Sprintf("%s contains in row %s", string(board[i][j]), string(i))] = true
				sb[fmt.Sprintf("%s contains in column %s", string(board[i][j]), string(j))] = true
			}
		}
	}

	return true
}