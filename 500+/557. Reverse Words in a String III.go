//
//date： 2017-6-19
//desc:  tired
//

func reverseWords(a string) string {
    
    left, right := 0, 0
    s  := []rune(a)
    length := len(s)
    v := 0

    for v < length {
        
        for v < length && s[v] == 32 {
            v++
        }
        
        if v == length { return string(s) }
        
        left = v
        
        
        for  v < length && s[v] != 32 {
            v++
        }
        
        if v == length {
            reverse(s,left,length - 1)
            return string(s)
        }
        
        right = v - 1
        reverse(s, left, right)

    }
    return string(s)
}

func reverse(s []rune, left, right int) {
    for left < right {
        a := s[left]
        s[left] = s[right]
        s[right] = a
        left++
        right--
    }
}
