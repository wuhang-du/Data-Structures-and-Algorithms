//
// index is big information
func findDuplicates(nums []int) []int {
	a := []int{}
	for k, v := range nums {
		for v != nums[v-1] {
			nums[k], nums[v-1] = nums[v-1], nums[k]
			v = nums[k]
		}

	}

	for k, v := range nums {
		if k != v-1 {
			a = append(a, v)
		}
	}

	return a
}
