//
//date:2017-6-1
//desc:smart go
//

func findRelativeRanks(nums []int) []string {
    length := len(nums)
    temp := make([]int,length)
    
    for i, v := range nums{
        temp[i] = v
    }
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    mymap := make(map[int]string,length)
    
    for i, v := range nums {
        switch i {
            case 0:
                mymap[v] = "Gold Medal"
            case 1:
                mymap[v] = "Silver Medal"
            case 2:
                mymap[v] = "Bronze Medal"
            default:
                mymap[v] =  fmt.Sprintf("%d",i+1)
        }
    }
    
    answer := make([]string,length)
    for i, v :=  range temp{
        answer[i] = mymap[v]
    }
    return answer
}
