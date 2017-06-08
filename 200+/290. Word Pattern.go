//
//date: 2017-6-8
//desc: 互相哈希。刚开始只是单向的hash， 这道题本质上是双向的。
//		做那道测试用例的时候，应该想到的，互相的逻辑是相似的。
//		所以，以后要开始先写测试用例，测试用例是思考的过程。
//

func wordPattern(pattern string, str string) bool {
    hash := make(map[rune]string)
    hash2 := make(map[string]bool)
    temp := strings.Split(str," ")
    length := len(pattern)
    if length != len(temp) { return false }
    
    for i, v := range pattern {
        if hash[v] == "" {
            if hash2[temp[i]] != false {
                return false
            }
            hash2[temp[i]] = true
            hash[v] = temp[i]
            
        } else {
            if temp[i] != hash[v] {
                return false
            }
        }
    }
 
    return true
}


func wordPattern(pattern string, str string) bool {
    temp := strings.Split(str," ")  
    array := make(map[byte][]int,1) 
    length := len(pattern)
    if length != len(temp) { return false }
    
    for i:= 0; i < length; i++ {    
        b , ok := array[pattern[i]]     
        if(!ok) {               
            b = make([]int, 0) 
        }
        b  = append(b,i)
       array[pattern[i]] = b
    }
    
    
    first := make([]int,0)

    for _, v := range array {
        
        prev := v[0]
        first = append(first,prev)
        for _, v2 := range v {
            //if  strings.Compare(temp[prev],temp[v2]) == -1 {
            if temp[prev] != temp[v2] {    
                return false
            }
        }
    }
    
    lenarray := len(first)
    
    for i := 0; i < lenarray; i++ {
        
        for j := i + 1; j < lenarray; j++ {
            if temp[first[i]] == temp[first[j]]  {
                return false
            }
        }
    }
    return true
    
}
