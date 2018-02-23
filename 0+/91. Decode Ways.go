// f(s) = n 表示 s时 为 n
import "strconv"

// f(s) = f(s[1:]) + f(s[2:]) if int(s[0:2]) <= 26

// f(s+2) = f(s+1) + f(s) s[2] <= 26

func numDecodings(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	first := 1
	second := 1

	for i := 1; i < len(s); i++ {

		b, _ := strconv.Atoi(s[i-1 : i+1])
		tmp := 0
		if b < 27 {
			tmp = second + first
		} else {
			tmp = second
		}
		first, second = second, tmp
	}
	return second
}
