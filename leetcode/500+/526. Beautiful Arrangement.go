func countArrangement(N int) int {
	nums := make([]int, N+1, N+1)

	for i := 0; i < N+1; i++ {
		nums[i] = i
	}
	count := 0
	generate(nums, 1, &count)
	return count
}

func generate(nums []int, pos int, count *int) {
	if pos == len(nums) {
		*count++
	}
	for i := pos; i < len(nums); i++ {
		nums[pos], nums[i] = nums[i], nums[pos]

		if nums[pos]%pos == 0 || pos%nums[pos] == 0 {
			generate(nums, pos+1, count)
		}
		nums[pos], nums[i] = nums[i], nums[pos]
	}
}
