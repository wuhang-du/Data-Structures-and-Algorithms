// date: 2017-7-24
// desc: new week
/*
    [2, 6, 4, 8, 10, 9, 15]   
    [1,2,3,8,7,6]  左边的界限是左边的数 都小于这个右边的所有的数 右边的界限是指 在右边的数都大于等于 该界限左边的数
    
    [1,3,2,2,5,6]  
 */

func findUnsortedSubarray(nums []int) int {
    
    left := 0
    for i := 0; i < len(nums) - 1; i++ {
        j := i + 1
        for j < len(nums) && nums[i] <= nums[j] { j++}
        if j != len(nums) { break }
        left++
        
    }

    right := len(nums) - 1 
    for i := len(nums) - 1 ; i >= 0; i-- {
        if i <= left {  return 0 }
        j := i - 1
        for j >= 0 && nums[i] >= nums[j] {
            j--
        }
        if j != -1 { break}
        right--
    }

    return right - left + 1
}
