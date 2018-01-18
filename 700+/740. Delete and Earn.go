func deleteAndEarn(nums []int) int {
	var cache = make([]int, 10001, 10001)
	answer := 0

	for _, v := range nums {
		cache[v] += v
	}
	maxN_1 := 0
	maxN_2 := 0

	for _, v := range cache {
		answer = v + maxN_2
		if answer < maxN_1 {
			answer = maxN_1
		}
		maxN_2 = maxN_1
		maxN_1 = answer
	}
	return answer
}

