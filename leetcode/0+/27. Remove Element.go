
// date: 2017-7-22
// desc: come on
//
/*
[1] 1  0
[1,2,3] 1 [3,2,1]
[1,2,3] 3 
[1,2,3] 4
*/


func removeElement(nums []int, val int) int {
    left := 0
    right := len(nums) - 1
    for left <= right {
        for left <= right && nums[left] != val {
            left++
        }
        for left <= right && nums[right] == val {
            right--
        }
        if left < right {
            nums[left], nums[right] = nums[right], nums[left]
            left++
            right--
        }
    }
    return left
}

func removeElement() {
	i := 0
	for j :=0; j < len(num); j++ {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return i
}
