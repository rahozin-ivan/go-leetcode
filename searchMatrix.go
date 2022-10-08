package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if len(row) == 0 {
			continue
		}
		if row[0] > target {
			return false
		}
		if row[len(row)-1] < target {
			continue
		}
		for _, col := range row {
			if col == target {
				return true
			}
		}
	}
	return false
}
