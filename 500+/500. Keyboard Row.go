// date:2017-7-31

func findWords(words []string) []string {
    hash := map[byte]int{'Q':1,'W':1,'E':1,'R':1,'T':1,'Y':1,'U':1,'I':1,'O':1,'P':1,
                         'A':2,'S':2,'D':2,'F':2,'G':2,'H':2,'J':2,'K':2,'L':2,
                         'Z':3,'X':3,'C':3,'V':3,'B':3,'N':3,'M':3,}
    var answer []string
    for _, v := range words {
        if v == "" {
            answer = append(answer, v)
            continue
        }
        t := []byte(strings.ToUpper(v))
        sign := hash[t[0]]
        for  _, m := range t {
            if hash[m] != sign {
                sign = hash[m]
                break
            }
        }     
        if sign == hash[t[0]] {
            answer = append(answer, v)
        }
    }
   return answer 
}
