//
// wrong memory use

func permute(nums []int) [][]int {
	switch len(nums) {
	case 0:
		return [][]int{{}}
	case 1:
		return [][]int{nums}
	default:
	}

	answer := make([][]int, 0, len(nums))
	for index, _ := range nums {
		nums[0], nums[index] = nums[index], nums[0]
		tmp := get(nums[0], nums[1:])
		for _, v := range tmp {
			answer = append(answer, v)
		}
		nums[0], nums[index] = nums[index], nums[0]
	}
	return answer
}

func get(start int, nums []int) [][]int {

	if len(nums) < 2 {
		return [][]int{{nums[0], start}}
	}

	answer := make([][]int, 0, len(nums))
	for index, _ := range nums {

		nums[0], nums[index] = nums[index], nums[0]
		tmp := get(nums[0], nums[1:])
		for _, v := range tmp {
			v = append(v, start)
			answer = append(answer, v)
		}
		nums[0], nums[index] = nums[index], nums[0]
	}
	return answer

}

// after update
type result [][]int

func permute(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{{}}
	}

	answer := new(result)
	answer.get(nums, 0)

	return *answer
}

func (r *result) get(nums []int, pos int) {
	if pos >= len(nums) {
		newSlice := make([]int, len(nums))
		copy(newSlice, nums)
		*r = append(*r, newSlice)
	}

	for index := pos; index < len(nums); index++ {
		nums[pos], nums[index] = nums[index], nums[pos]
		r.get(nums, pos+1)
		nums[pos], nums[index] = nums[index], nums[pos]
	}
}
