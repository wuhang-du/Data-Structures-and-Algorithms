func minPathSum(grid [][]int) int {
	tmp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[i]); j++ {
			// now (i,j) is its postion
			// check i-1 && j-1
			n := 0
			if i-1 >= 0 {
				n = tmp[j]
			}
			if j-1 >= 0 {
				if n == 0 {
					n = grid[i][j-1]
				} else if n > grid[i][j-1] {
					n = grid[i][j-1]
				}
			}
			grid[i][j] += n
			tmp[j] = grid[i][j]
		}
	}
	return tmp[len(tmp)-1]
}
