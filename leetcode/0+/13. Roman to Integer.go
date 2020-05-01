//
//date: 2017-6-15
//desc: easy
//


// I(1)、V(5)、X(10)、l(50)、C(100)、D(500)和M(1000)

func romanToInt(s string) int {
    
    m := map[byte]int{'I':1,'V':5,'X':10,'L':50,'C':100,'D':500,'M':1000}
    
    answer := 0
    length := len(s) - 1
    
    for i := 0; i < length; i++ {
        if m[s[i]] < m[s[i+1]] {
            answer -= m[s[i]]
        } else {
            answer += m[s[i]]
        }
    }
   
    answer += m[s[length]]
    return answer
}
