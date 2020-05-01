//
//date: 2017-6-20
//desc: 快排。lo值是第一个大于等于key值的位置，
// 所以： 当两面夹击的时候 条件是 lo <= high 
		  当单次循环的时候 条件是 lo <= right - 1 即 lo < right

// 排序 隔开找数即可。
// 1，4，3，2
func arrayPairSum(nums []int) int {
    
    quickSort(nums, 0 ,len(nums) - 1)
    sum := 0
    for i := 0; i < len(nums); i += 2{
           sum += nums[i]
    }
    return sum
}
func partition(a []int, lo, hi int) int {
    key := a[hi]
    
    for i := lo; i < hi ; i++ {
        if a[i] < key {
            a[lo], a[i] = a[i], a[lo]
            lo++
        }
    }
    a[lo], a[hi] = a[hi], a[lo]
    return lo
    
    
    /*
    right := hi
    hi--
    for lo <= hi {
        for a[lo] < a[right] { lo++ }
        for hi >= 0 &&  a[hi] >= a[right] { hi-- }
        if lo < hi {
            a[lo] , a[hi] = a[hi], a[lo]
        }
    }
    a[lo], a[right] = a[right], a[lo]
    return lo
    */
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
