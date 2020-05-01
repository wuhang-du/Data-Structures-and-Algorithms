func subsets(nums []int) [][]int {
	answer := [][]int{[]int{}}
	buf := make([]int, len(nums))
	get(nums, &buf, &answer, 0)
	return answer
}

func get(nums []int, prev *[]int, a *[][]int, pos int) {
	for i := 0; i < len(nums); i++ {
		(*prev)[pos] = nums[i]
		tmp := make([]int, pos+1)
		copy(tmp, (*prev)[:pos+1])
		*a = append(*a, tmp)
		get(nums[i+1:], prev, a, pos+1)
	}
}
