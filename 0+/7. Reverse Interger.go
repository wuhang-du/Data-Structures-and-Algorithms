//
//date: 2017-6-7
//desc: two methods, I love the last, though it's not best.
//

const UINT32_MAX = ^uint32(0)
const INT32_MAX = int(UINT32_MAX >> 1)
const INT32_MIN = - INT32_MAX - 1

func reverse(x int) int {
    answer := 0
    
    for x != 0 {
        answer = answer * 10 + x % 10
        if(answer > INT32_MAX || answer < INT32_MIN) {
            return 0
        }
        x /= 10
    }
    return answer
}

/*
func reverse(x int) int {
    sign := false  //默认大于等于0
    if( x < 0) {
        sign = true
        x = -x
    }
    
    array := make([]int, 10)
    
    for i := 0; x > 0; i++ {
        array[i] = x % 10
        x = x / 10
    }
    
    start, end  := 0, 0
    
    for i := 0 ; i < len(array); i++ {
        if( array[i] != 0) {
            start = i
            break
        }
    }
    
    for j := len(array) - 1; j > -1; j-- {
        if( array[j] != 0) {
            end = j
            break
        }
    }
    
    temp := array[start:end + 1]
    
    answer := 0
    for i, v := range temp {
        switch {
            case i < 9:
                answer = answer * 10 + v
            case i == 9:
                if( answer > 214748364) {
                    return 0 //已经溢出
                }
                
                if( sign && v > 8){
                    return 0
                }
                
                if(sign == false && v > 7) {
                    return 0
                }
                
                answer = answer * 10 + v
        }
    }
    if(sign) { return -answer }
    return answer
}
*/
