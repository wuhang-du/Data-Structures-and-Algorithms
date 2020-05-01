//
// date: 2018-1-22
//
func rangeBitwiseAnd(m int, n int) int {
	answer := 0
	for i := uint(31); i > 0; i-- {
		tmp := 1 << (i - 1)
		if m&tmp != n&tmp {
			return answer
		}
		answer |= m & tmp
	}
	return answer
}
