package main

//剑指 Offer 04. 二维数组中的查找

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	xx := len(matrix)
	yy := len(matrix[0])
	x := 0
	y := yy - 1

	for x >= 0 && x < xx && y >= 0 && y < yy {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else if matrix[x][y] > target {
			y--
		}

	}
	return false
}
