// first
// 10^9----2^32 32/4 = 8 byte
// T:O(n^2)
// S:O(1)
var cache = make(map[int]int)

func totalHammingDistance(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	answer := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			answer += get(nums[i] ^ nums[j])
		}
	}
	return answer
}

func get(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	temp := n
	n = (n & 0x5555555555555555) + ((n >> 1) & 0x5555555555555555)
	n = (n & 0x3333333333333333) + ((n >> 2) & 0x3333333333333333)
	n = (n & 0x0F0F0F0F0F0F0F0F) + ((n >> 4) & 0x0F0F0F0F0F0F0F0F)
	n = (n & 0x00FF00FF00FF00FF) + ((n >> 8) & 0x00FF00FF00FF00FF)
	n = (n & 0x0000FFFF0000FFFF) + ((n >> 16) & 0x0000FFFF0000FFFF)
	n = (n & 0x00000000FFFFFFFF) + ((n >> 32) & 0x00000000FFFFFFFF)

	cache[temp] = n
	return n
}

// second

// 10^9----2^32 32/4 = 8 byte
// T:O(N)
// S:O(1)

func totalHammingDistance(nums []int) int {
	total := 0
	length := len(nums)

	for i := 0; i < 32; i++ {
		p := 0
		for j := 0; j < length; j++ {
			p += (nums[j] >> uint(i)) & 1
		}
		total += p * (length - p)
	}
	return total
}
