//
//date: 2017-6-19
//desc: easy
//

func fizzBuzz(n int) []string {
    answer := make([]string,n)
    for i := 1; i <= n; i++ {
        var temp string
        if i % 3 == 0 {
            temp = "Fizz" 
        }
        if i % 5 == 0 {
            temp += "Buzz"
        }
        if temp == "" {
            temp = strconv.Itoa(i)
        }
        answer[i - 1] = temp
    }
    return answer
}
