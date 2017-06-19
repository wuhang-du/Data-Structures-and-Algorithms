//
//date: 2017-6-19
//desc: array
//

func canConstruct(ransomNote string, magazine string) bool {
    var check [26]int
    for _, v := range magazine {
        check[v - 'a']++
    }
    for _, v := range ransomNote {
        check[v - 'a']--
        if check[v - 'a'] < 0 {
            return false
        }
    }
    return true
}
