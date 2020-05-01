// date:2017-7-26
// desc: come on
/*
C(n+1,i)=C(n,i)+C(n,i-1)
0     1
1    1 1
2   1 2 1
3  1 3 3 1
4 1 4 6 4 1

生成一半，另一半镜像

*/

func getRow(rowIndex int) []int {
    
    if rowIndex <= 0 {
        return []int{1}
    }
    answer := make([]int, rowIndex + 1)
    answer[0] = 1
 
    //生成前一半的数组
    for i := 1; i <= rowIndex ; i++ {
        temp := 0
        half := (i + 2) / 2
        for j := 0; j < half; j++ {
            temp2 := answer[j]
            answer[j] += temp
            temp = temp2
        }
        answer[half] = answer[half - 1]
    }
    
    
    left := 0
    right := rowIndex 
    for left < right {
        answer[right] = answer[left]
        left++
        right--
    }
    
    return answer     
}
