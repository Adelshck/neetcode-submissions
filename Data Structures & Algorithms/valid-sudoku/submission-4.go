func isValidSudoku(board [][]byte) bool {
    var rows, cols, sqs [9]uint16

    for i := range 9 {
        for j := range 9 {
            val := board[i][j]
            if val == '.' {
                continue
            }

            bit := uint16(1 << (val - '1'))
            idx := (i/3)*3 + (j/3)

            if (rows[i]&bit) != 0 || (cols[j]&bit) != 0 || (sqs[idx]&bit) != 0 {
                return false
            }

            rows[i] |= bit
            cols[j] |= bit
            sqs[idx] |= bit
        }
    }
    return true
}
