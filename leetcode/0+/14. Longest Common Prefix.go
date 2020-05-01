// date: 2017-8-1
// desc: come on

func longestCommonPrefix(strs []string) string {
    length := len(strs) 
    if length == 0 {
        return ""
    }
    
    temp := []byte(strs[length / 2])
    check := 10000     
    for _, str := range strs {
        v := []byte(str)
        i, j := 0, 0
        for ; i < check && i < len(temp) && j < len(v); {
            if temp[i] != v[j] {
                break;
            }
            i++
            j++
        }
        check = i
    }
    return string(temp[0:check])
}
