//
//date: 2017-6-16
//desc：在刚开始用哈希做了实现之后，就在想能不能连续初始化hash键，其实就是数组了。哈哈。被自己蠢哭。还有就是make可以指大小。
//


func isAnagram(s string, t string) bool {
    if len(s) != len(t) { return false }
    
    var temp [26]int
    
    for _, v := range s {
        temp[v - 'a']++
    } 
    
    for _, v := range t {
        temp[v - 'a']--
        if temp[v - 'a'] < 0 {
            return false
        }
    }
    
    return true
    
    /*
    length := len(s)
    if length != len(t) { return false }
    temp := make(map[byte]int,26)
    for i := 0; i < length; i++ {
        temp[s[i]]++
        temp[t[i]]--
    }
    
    for _, v := range  temp {
        if v > 0 {
            return false
        }
    }
    
    return true
    */
}
