func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	sqs := make(map[[2]int]map[byte]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			if rows[i] == nil {
				rows[i] = make(map[byte]bool)
			}
			if cols[j] == nil {
				cols[j] = make(map[byte]bool)
			}

			if sqs[[2]int{i / 3, j / 3}] == nil {
				sqs[[2]int{i / 3, j / 3}] = make(map[byte]bool)
			}
			if _, exist := rows[i][val]; exist {
				return false
			}
			rows[i][val] = true

			if _, exist := cols[j][val]; exist {
				return false
			}
			cols[j][val] = true

			if _, exist := sqs[[2]int{i / 3, j / 3}][val]; exist {
				return false
			}
			sqs[[2]int{i / 3, j / 3}][val] = true
		}
	}
	return true
}
