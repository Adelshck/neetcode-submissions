func searchMatrix(matrix [][]int, target int) bool {
	lRow := 0
	lCol := 0
	rRow := len(matrix) - 1
	rCol := len(matrix[0]) - 1



	for lRow <=rRow  {
		midRow := lRow + (rRow - lRow) / 2
		midCol := lCol + (rCol - lCol) / 2
		if matrix[midRow][0] <= target && matrix[midRow][len(matrix[0]) - 1] >= target {
				for lCol <= rCol {
					midCol = lCol + (rCol - lCol) / 2
					if matrix[midRow][midCol] > target {
						rCol = midCol - 1
					} else if matrix[midRow][midCol] < target {
						lCol = midCol + 1
					} else {
						return true
					}
				}
				return false
			}
		
		if matrix[midRow][midCol] > target {
			rRow = midRow - 1
		} else if matrix[midRow][midCol] < target {
			lRow = midRow + 1
		}
		}
	return false
}