//
//date： 2017-7-10
//desc:  array， map, or == ==
//

func reverseVowels(s string) string {
    left := 0
    right := len(s) - 1
    index := map[byte]int{'A':1,'a':1,'E':1,'e':1,'I':1,'i':1,'O':1,'o':1,'U':1,'u':1}
    s1 := []byte(s)
    for left < right {
        for left < right && index[s1[left]] != 1 {
            left++
        }
        for right > 0 && index[s1[right]] != 1 {
            right--
        }
        if left < right {
            s1[left], s1[right] = s1[right], s1[left]
        }
        left++
        right--
    }
    return string(s1)
}
