//
//date: 2017-6-1
//desc: easy
//

func majorityElement(nums []int) int {
    length := len(nums) >> 1 + 1
    answer := make(map[int]int, length);
    for _, v := range nums {
        answer[v]++
    }
    for i, v := range answer {
        if( v >= length){
            return i
        }
    }
    return nums[0]
}
