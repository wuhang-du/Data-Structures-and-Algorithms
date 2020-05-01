//
//date: 2017-6-7
//desc: boring
//

func isPalindrome(s string) bool {
    s = strings.ToUpper(s)
    length := len(s)
    start, end := 0, length - 1
    
    for ; start < end ; {
        
        for ; start < length && (s[start] < 48 || s[start] > 90 || (s[start] > 57 && s[start] < 65)); {
            start++
        }
        for ; end > -1 && (s[end] < 48 || s[end] > 90 || (s[end] > 57 && s[end] < 65)); {
            end--
        } 
        if start < end {
            if s[start] != s[end] {
                return false
            }
            start++
            end--
        }
    }
    return true
}
