//
//date: 2017-6-6
//desc: no words to say
//

func findNthDigit(n int) int {
    if n < 10 { return n }
    ten := 1
    max := 9
    for i := 2; i < 10; i++ {
        ten = ten * 10
        temp := max
        max += ten * i * 9
        if n < max {
            n -= temp   // 代表这是i位里面的第N个。
            temp = ten + (n - 1) / i
            num := n % i
            var array []int
            for ; temp > 0 ; {
                array = append(array, temp % 10)
                temp = temp / 10
            }
            if num == 0 { 
                return array[0] 
            } else { 
                return array[i - num] 
            }
        }
    }
    return 0
}
