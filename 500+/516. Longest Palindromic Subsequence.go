// f(i)的回文个数
// 加上新的i 之后
// dp[i][j] = dp[i+1][j-1] + 2 if s.charAt(i) == s.charAt(j)
// otherwise, dp[i][j] = Math.max(dp[i+1][j], dp[i][j-1])
// Initialization: dp[i][i] = 1
//

// first edition
func longestPalindromeSubseq(s string) int {
	return check([]byte(s))
}

func check(s []byte) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if s[0] == s[len(s)-1] {
		return 2 + check(s[1:len(s)-1])
	} else {
		a := check(s[0 : len(s)-1])
		b := check(s[1:len(s)])
		if b > a {
			a = b
		}
		return a
	}
}

// f(i)的回文个数
// 加上新的i 之后
// dp[i][j] = dp[i+1][j-1] + 2 if s.charAt(i) == s.charAt(j)
// otherwise, dp[i][j] = Math.max(dp[i+1][j], dp[i][j-1])
// Initialization: dp[i][i] = 1
//

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	cur := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		cur[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				cur[j] = dp[j-1] + 2
			} else {
				cur[j] = max(cur[j-1], dp[j])
			}
		}
		dp, cur = cur, dp
	}
	return dp[n-1]
}
