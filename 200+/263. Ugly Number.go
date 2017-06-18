//
//date: 2017-6-18
//desc：stupid
//

for n % 2 == 0 {
	n = n / 2
}


func isUgly(num int) bool {
    if num <= 0 {return false}
    if num < 4  {return true}
    
    if( num % 2 != 0 && num % 3 != 0 && num % 5 != 0) {
        return false
    }
    
    now := 3
    temp := make([]int,0)
    
    temp = append(temp, 1)
    temp = append(temp, 2)
    temp = append(temp, 3)
    
    index := 0
    
    temp1 := make([]int, 100)
    
    for {
        check := temp[index] * 5
        
        j := 0
        
        for i := index ; i < len(temp); i = i+1 {
            
            m := now / temp[i]
            switch {
            case m < 2:
                m = 2
            case m < 3:
                m = 3
            default:
                m = 5
            }
            
            
            temp1[j] = temp[i] * m
            j++
            if m == 2 {  break }
        } 

        
        min := temp1[0]
        
        for i := 1; i < j; i++ {
            if min > temp1[i] {
                min = temp1[i]
            }
        }
    
        if min == check { index++ }
        if min == num { return true }
        if min > num { return false }
        
        now = min
        temp = append(temp, min)
    }
}

