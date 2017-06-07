//
//date: 2017-6-7
//desc: come on
//

func trailingZeroes(n int) int {
    if(n <= 0) { return 0 }
    count := 0
    for num := n/5; num > 0; {
        count += num
        num /= 5
    }
    return count
}
