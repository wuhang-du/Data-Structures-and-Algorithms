func findRepeatedDnaSequences(s string) []string {

	record := map[string]int{}
	answer := []string{}

	for i := 0; i < len(s)-9; i++ {
		if n, ok := record[s[i:i+10]]; ok && n == 1 {
			answer = append(answer, s[i:i+10])
		}
		record[s[i:i+10]]++
	}
	return answer
}
