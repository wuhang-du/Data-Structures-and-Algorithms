// 已经排序 考虑 二分查找。
// 有一个问题：抛掉字母论。 以 1，2，3，4，5，6 为主。
// target 为 3.
// 如果给的是 4 5 6 相关，就找大于 target 的。
// 如果找不到，就找最小的。也就是 最左边的那个数了。

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target || letters[0] > target {
		return letters[0]
	}

	left := 0
	right := len(letters) - 1
	target++
	for right-left > 1 {
		mid := (right + left) >> 1
		if letters[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return letters[right]
}
