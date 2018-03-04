//
// 之前考虑的是 递归或者动态规划一样的，比较耗内存。
// 现在考虑的是 如何只使用 k 空间的内存，去做这个事情

func combine(n int, k int) [][]int {
	answer := [][]int{}
	tmp := make([]int, k)
	create(1, n, k, 0, &tmp, &answer)
	return answer
}

func create(start int, end int, k int, pos int, tmp *[]int, answer *[][]int) {

	if k == 0 {
		t := make([]int, len(*tmp))
		copy(t, *tmp)
		*answer = append(*answer, t)
		return
	}

	for i := start; i <= end+1-k; i++ {
		(*tmp)[pos] = i
		create(i+1, end, k-1, pos+1, tmp, answer)
	}
}
