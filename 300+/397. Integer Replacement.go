func integerReplacement(n int) int {
	count := 0

	for n > 1 {
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || ((n>>1)&1 == 0) {
			n -= 1
		} else {
			n += 1
		}
		count++
	}

	return count
}
