//
//date: 2017-6-16
//desc: 为了不使用取余的操作，将 赋值裂开了两个for 循环。
//

func rotate(nums []int, k int)  {
    
    length := len(nums)
    k = k % length
    
    temp := make([]int, length)
    count := 0
    
    for i := length - k; i < length; i++ {
        temp[count] = nums[i]
        count++
    } 
    
    for m := 0 ; m < length - k ; m++ {
        temp[count] = nums[m]
        count++
    }
    
    for k, v := range temp {
        nums[k] = v
    }
    
}
