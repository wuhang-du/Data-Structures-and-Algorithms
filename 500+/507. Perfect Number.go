// date: 2017-7-18
// desc: new method
//
//  28: 1 2 4 7 14 28
//  10000 ： 5000 2
//  正向： 算出所有的公约数，求和
//  反向： 10000 = 5000 * 2 = 2500 * 4 = 1250 *8 = 625 * 16 = 125 * 80 = 25 * 400 = 5 * 2000 = 1 * 10000
//         28 = 14 * 2 = 7 * 4 = 1 * 28
//          9 = 3 * 3 = 1 * 9
//          3 = 3 * 2 //不整除，除数就增加
//  大素数 解决办法： 32582657
func checkPerfectNumber(num int) bool {
    if num < 3 {
        return false
    }
    temp := num - 1
    devide := 2
    end := int(math.Sqrt(float64(num)))
    
    for num != devide {
        a := 0
        for devide <= end {
            if num % devide == 0 {
                temp -= devide
                a = num / devide
                temp -= a
                //fmt.Printf("%d %d \n",a,devide)
                break
            }
            devide++
        }
        
        if devide > end {
            return false
        }
        b := 2
        for b <= a {
            if a % b == 0 {
                devide = devide * b
                break
            }
            b++
        } 
    }
    
   return temp == 0

}
