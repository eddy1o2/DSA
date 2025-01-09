package array

// IsValidSudoku https://leetcode.com/problems/valid-sudoku/
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			b := board[i][j]
			// Multiple with `3` because we flat the idx of square to a slice from 0-9
			squareIdx := (i/3)*3 + j/3
			if rows[i][b] || columns[j][b] || squares[squareIdx][b] {
				return false
			}
			rows[i][b] = true
			columns[j][b] = true
			squares[squareIdx][b] = true
		}
	}
	return true
}
