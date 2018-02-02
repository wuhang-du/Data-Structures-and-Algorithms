func PredictTheWinner(nums []int) bool {
	return get(nums, 0, 0, true)
}

func get(nums []int, left, right int, sign bool) bool {

	if len(nums) == 0 {
		return left >= right
	}
	if sign {
		return get(nums[1:], left+nums[0], right, !sign) || get(nums[:len(nums)-1], left+nums[len(nums)-1], right, !sign)
	} else {
		return get(nums[1:], left, right+nums[0], !sign) && get(nums[:len(nums)-1], left, right+nums[len(nums)-1], !sign)
	}
}
