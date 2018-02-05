// O(n2)
// O(n)
func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	tmp := make([]int, m)
	tmp[0] = 1

	for i := 0; i < n; i++ {
		b := 0
		for j := 0; j < m; j++ {
			tmp[j] += b
			b = tmp[j]
		}
	}
	return tmp[m-1]
}
