//
//date: 2017-6-9
//desc: time is too long
//

func isIsomorphic(s string, t string) bool {
    hashs := make(map[byte]int)
    hasht := make(map[byte]int)
    count := len(s)
    if count != len(t) { return false}
    for i:= 0; i < count ; i++ {
        a , b := hashs[s[i]] 
        c , d := hasht[t[i]]
        if  a != c || b != d { return false} 
        hashs[s[i]] = i
        hasht[t[i]] = i
    }
    return true
}
