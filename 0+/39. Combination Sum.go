// 这道题的关键有以下几点： 对数组的排序
// 在排序的基础上
// 关于处理重复元素就有了新的点，那就是 c[i:]
// 它是一个逐渐递增的过程
// 不会同时出现 2 2 3 和 2 3 2 这两种情况。

import "sort"

func combinationSum(c []int, t int) [][]int {
	a := [][]int{}
	tmp := []int{}
	sort.Ints(c)
	get(c, t, &tmp, &a)
	return a
}

func get(c []int, t int, tmp *[]int, a *[][]int) {

	if t == 0 {
		m := make([]int, len(*tmp))
		copy(m, *tmp)
		*a = append(*a, m)
		return
	}

	for i := 0; i < len(c) && c[i] <= t; i++ {
		*tmp = append(*tmp, c[i])
		get(c[i:], t-c[i], tmp, a)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
