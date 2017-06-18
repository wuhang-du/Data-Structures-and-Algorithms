//
//date: 2017-6-18
//desc：上一次是发现有重复的代码，则有优化的途径。这一次，是发现有重复的逻辑运算，则有优化的途径。第一次 100% beats。
//

func nthUglyNumber(n int) int {
    temp := new([1690]int)
    temp[0] = 1
    count := 1
    two, three, five := 0, 0, 0
    
    for count < n {
        two_now := temp[two] * 2
        three_now := temp[three] * 3
        five_now := temp[five] * 5
        min := two_now
        if min > three_now { min = three_now}
        if min > five_now { min = five_now }
        
        temp[count] = min
        count++
        if min == two_now { two++}
        if min == three_now { three++ }
        if min == five_now  { five++ }
    }
    
    return temp[count - 1]
}
