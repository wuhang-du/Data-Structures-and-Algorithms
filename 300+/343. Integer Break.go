/* date:2017-8-16
    f(n) 状态：表示f(n)此时能够分解的值
    f(2) = 2
    f(3) = 2

    f(4) 3,f(3)  2 f(2) * 2 f(2)   4
    f(5) 4,f(4)  3 f(3) * 2 f(2)   6
    f(6) 5,f(5)  4 f(4) * 2 f(2)   3 f(3) * 3 f(3)    9
    f(7) 6,f(6)  5 f(5) * 2 f(2)   4 f(4) * 3 f(3)    12
    f(8)         f(6) *  f(2)   f(5) * f(3)  f(4) * f(4)    18
                              
   
   f(9)         f(7) * f(2)   f(6) * f(3)    f(5) * f(4)    27
      
    f(10)        8 f(8) * 2 f(2)   7 f(7) * 3 f(3)   6 f(6) * 4 f(4)   5  f(5) * 5 f(5)
    
                  f(8) * f(2)   f(7) * f(3)  f(6) * f(4)   f(5) * f(5)
                  
    f(11)        27*2=54  36 48 6*9=54
    f(n) > n     
*/


func integerBreak(n int) int {
    if n < 4 { return n-1}
    var temp [58]int
    temp[2] = 2
    temp[3] = 3
    max := 0

    for i := 4; i <= n; i++ {
        max = 0
        for j:= i -2; j >= 2; j-- {
            k := temp[j] * temp[i-j]
            if k > max { max = k}
        }
        temp[i] = max
    }
    return max
}
