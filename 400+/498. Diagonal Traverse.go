func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	x, y := 0, 0
	row, line := len(matrix), len(matrix[0])
	sign := true
	a := make([]int, row*line)
	pos := 0
	for x < row && y < line {
		a[pos] = matrix[x][y]
		pos++
		if sign {
			// 正常斜向上
			if x-1 > -1 && y+1 < line {
				x = x - 1
				y = y + 1
				continue
			}

			if y+1 < line {
				y = y + 1
				sign = false
				continue
			}

			if x+1 < row {
				x = x + 1
				sign = false
				continue
			}
			break
		} else {
			if x+1 < row && y-1 > -1 {
				x = x + 1
				y = y - 1
				continue
			}

			if x+1 < row {
				x = x + 1
				sign = true
				continue
			}
			if y+1 < line {
				y = y + 1
				sign = true
				continue
			}
			break
		}
	}
	return a
}
