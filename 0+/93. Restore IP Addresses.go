import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	tmp := make([]string, 4)
	answer := []string{}
	get(s, &tmp, 0, &answer)
	return answer
}

func get(s string, tmp *[]string, count int, answer *[]string) {
	if len(s) < 4-count && len(s) > (4-count)*3 {
		return
	}

	if count == 3 {
		if checknum(s) {
			(*tmp)[count] = s
			*answer = append(*answer, strings.Join(*tmp, "."))
		}
		return
	}

	for i := 1; i < 4 && i < len(s); i++ {
		if checknum(s[0:i]) {
			(*tmp)[count] = s[0:i]
			get(s[i:], tmp, count+1, answer)
		}
	}
}

func checknum(s string) bool {
	b, _ := strconv.Atoi(s)
	switch len(s) {
	case 1:
		return true
	case 2:
		return b > 9
	case 3:
		return b > 99 && b < 256
	default:
	}
	return false
}
