//
// substring 不包括跳动的字符。肯定是相邻的。
//
// 区别在于缓存的表现。我用的map, 人家用的[][]bool。
// 节省了很多的速度。
//

func countSubstrings(s string) int {
	//cache := map[string]bool{}
	count := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (i-j < 3 || dp[i-1][j+1]) {
				count++
				dp[i][j] = true
			}

		}
	}
	return count
}

func countSubstrings(s string) int {
	cache := map[string]bool{}
	count := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] == s[i] {

				if j+1 == i || j+2 == i {
					count++
					cache[s[j:i+1]] = true
					continue
				}

				if _, ok := cache[string(s[j+1:i])]; ok {
					count++
					cache[s[j:i+1]] = true
				}
			}
		}
		// 本次的单个字符
		count++
	}
	return count
}
