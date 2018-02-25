func findTargetSumWays(nums []int, S int) int {
	length := len(nums)
	count := 1 << uint32(length)

	t1 := make([]int, count)
	t2 := make([]int, count)

	l := 1
	for i := 0; i < length; i++ {

		for j := 0; j < l; j++ {
			// +
			t2[j] = t1[j] + nums[i]
			// -
			t2[j+l] = t1[j] - nums[i]
		}
		l += l
		t1, t2 = t2, t1
		//fmt.Println(t1)
	}
	// t1 保存着 新鲜的结果

	answer := 0
	for _, v := range t1 {
		if v == S {
			answer++
		}
	}

	return answer
}

func findTargetSumWays(nums []int, S int) int {
	return get(nums, 0, S)
}

func get(nums []int, i int, sum int) int {
	if len(nums) == i {
		if sum == 0 {
			return 1
		} else {
			return 0
		}
	}
	return get(nums, i+1, sum+nums[i]) + get(nums, i+1, sum-nums[i])
}
