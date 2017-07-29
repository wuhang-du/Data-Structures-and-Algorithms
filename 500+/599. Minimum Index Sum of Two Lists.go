// date: 2017-7-27
// desc: come on
/*
 [1,2,3,4,5] 
 [3,5,7]
 
 [md5(key)] index 选长度小的那个 做这个。
*/

func findRestaurant(list1 []string, list2 []string) []string {
    p1 := &list1
    p2 := &list2
    if len(list1) > len(list2) {
        p1 = &list2
        p2 = &list1
    }
    
    hash := make(map[string]int, len((*p1))) 
    for k, v := range *p1 {
        hash[v] = k
    }
    
    min := len(list1) + len(list2)
    var answer []string
    
    for k, v := range *p2 {
        if k > min {
            break
        }
        if i, ok := hash[v]; ok {
            temp := i + k
            if temp == min {
                answer = append(answer,v)
            }
            if temp < min {
                min = temp
                answer = (*p2)[k:k+1]
            }
        }
    }
    return answer
}
