//
// 贪心算法： for 里面的 if 比较就是 贪心法选择的标准。 这个标准很重要。
//
//设有n个正整数，将它们连接成一排，组成一个最大的多位整数。
//例如：n=3时，3个整数13，312，343，连成的最大整数为34331213。
//又如：n=4时，4个整数7，13，4，246，连成的最大整数为7424613
//那么 12, 121 呢?
//另外一个就是  拥有最优解，并且 最优解 可以由 小范围的 最优解 组成。

// 动态规划： 不到最后一步，不能知道到底哪一个是最好的。
// 状态转移方程。 从一个到另一个。

// 贪心算法仅仅是动态规划的一个平凡态。
// 多动脑子，就能少一些耗费。

func canJump(nums []int) bool {

	i := 0
	for i < len(nums) {

		prev := i
		if i+nums[i] >= len(nums)-1 {
			break
		}

		tmp := 0
		next := i

		for j := prev + nums[prev]; i <= j; {
			if i+nums[i] > tmp {
				tmp = i + nums[i]
				next = i
			}
			i++
		}

		if next == prev {
			return false
		}
		i = next
	}
	return true
}
