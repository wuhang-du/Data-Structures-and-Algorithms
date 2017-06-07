//
//date: 2017-6-7
//desc: come on
//


func removeDuplicates(nums []int) int {
    if(len(nums) == 0) { return 0 }
    i := 0
    
    for _, v := range nums {
        if(nums[i] != v){
            i++
            nums[i] = v
        }
    }
    return i + 1
}
