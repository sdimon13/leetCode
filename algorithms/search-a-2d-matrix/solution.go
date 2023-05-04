package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	rowsSize := len(matrix)
	colsSize := len(matrix[0])

	left := 0
	right := rowsSize*colsSize - 1

	for left <= right {
		middle := (left + right) / 2
		current := matrix[middle/colsSize][middle%colsSize]

		if current > target {
			right = middle - 1
		} else if current < target {
			left = middle + 1
		} else {
			return true
		}
	}

	return false
}
