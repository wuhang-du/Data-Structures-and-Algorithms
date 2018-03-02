func findPeakElement(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	return buildHeap(nums)
}

func buildHeap(nums []int) int {
	n := len(nums)
	record := make(map[int]int)

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {

		j := i*2 + 1

		if j+1 < n && nums[j+1] > nums[j] {
			j++
		}

		if nums[i] < nums[j] {
			if _, ok := record[nums[j]]; !ok {
				record[nums[j]] = j
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if i, ok := record[nums[0]]; ok {
		return i
	}
	return 0
}
