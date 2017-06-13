//
//date: 2017-6-13
//desc: easy
//


func searchInsert(nums []int, target int) int {
    for k, v := range nums {
        if target <= v {
            return k
        }
    }
    return len(nums)
}

func searchInsert(nums []int, target int) int {
    if( nums[0] >= target ) { return 0}
    left, right := 0, len(nums) - 1
    if( nums[right] < target ) { return len(nums)}
    
    for right - left > 1 {
        middle := (left + right + 1) >> 1
        if nums[middle] >= target {
            right = middle
        } else {
            left = middle
        }
    }
    return right 
}
