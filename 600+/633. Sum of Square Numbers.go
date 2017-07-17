//
// date: 2017-7-17
// desc: new method
//
// think: non-negative 
// 0 1 2 3 4   5 6   7 8  9
// 0 1 4 9 16 25     36 49 64 81
//  a^2 = c - b^2 
//  26 = 5.09^2 
//  

func judgeSquareSum(c int) bool {
    right := int(math.Sqrt(float64(c)))
    left := 0 
    temp := 0
    for left <= right {
        temp = left * left + right * right
        switch {
            case temp == c:
                return true
            case temp > c:
                right--
            case temp < c:
                left++
        }
    }
    return false
}
