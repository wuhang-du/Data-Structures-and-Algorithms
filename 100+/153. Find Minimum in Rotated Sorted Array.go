func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	if nums[left] < nums[right] {
		return nums[left]
	}

	for right-left > 1 {
		mid := (left + right + 1) >> 1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid
		}
	}

	return nums[right]
}
