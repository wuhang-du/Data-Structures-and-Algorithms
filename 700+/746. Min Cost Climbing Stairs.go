// f(i) = min(f(i-1),f(i-2)) + a[i]
//
func minCostClimbingStairs(cost []int) int {

	first := 0  // i-2
	second := 0 // i-1

	for i := 0; i < len(cost); i++ {
		min := second
		if first < second {
			min = first
		}
		first, second = second, cost[i]+min
	}

	if first < second {
		return first
	} else {
		return second
	}
}
