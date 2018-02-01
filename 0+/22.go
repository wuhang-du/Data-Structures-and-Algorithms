func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}
	t := make([]byte, n*2)
	a := []string{}
	get(&t, n, n, &a, 0)
	return a
}

func get(t *[]byte, left, right int, a *[]string, i int) {

	if left == right {
		if left > 0 {
			(*t)[i] = '('
			get(t, left-1, right, a, i+1)
		} else {
			*a = append(*a, string(*t))
			return
		}

	}
	if left < right {
		(*t)[i] = ')'
		if left > 0 {
			get(t, left, right-1, a, i+1)
			(*t)[i] = '('
			get(t, left-1, right, a, i+1)
		} else {
			get(t, left, right-1, a, i+1)
		}
	}
}
