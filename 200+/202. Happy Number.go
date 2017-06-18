//
//date: 2017-6-18
//desc: 找规律，hash到 单个数
//

func isHappy(n int) bool {

    for n != 4 {
        temp := 0
        for n > 0 {
            m := n % 10
            temp += m * m
            n /= 10
        }
        if temp == 1 { return true}
        n = temp
    }
    
    return false
}
