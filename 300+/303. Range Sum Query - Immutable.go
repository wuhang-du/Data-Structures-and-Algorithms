//
//date: 2017-6-5
//desc: didn't get point
//

type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{nums}
}


func (this *NumArray) SumRange(i int, j int) int {
    sum := 0
    for index := i; index <= j; index++ {
        sum += this.nums[index]
    }
    return sum
}

// new method
type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    temp := make([]int, len(nums))
    index := 0
    for i, v := range nums {
        temp[i] = index + v
        index = temp[i]
    }
    return NumArray{temp}
}


func (this *NumArray) SumRange(i int, j int) int {
    switch {
        case i > j:
            return 0
        case i == 0:
            return this.nums[j]
        case i <= j:
            return this.nums[j] - this.nums[i - 1]
    }
    return 0
}
