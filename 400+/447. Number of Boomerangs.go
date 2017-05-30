//
//date: 2017-5-30
//desc: go begin
//


func numberOfBoomerangs(points [][]int) int {
    answer := 0
    for i, v := range points {
        mymap := make(map[int]int)
        for ii, vv := range points {
            if(ii != i) {
                key := ((v[0]-vv[0]) * (v[0]-vv[0]) + (v[1]-vv[1])*(v[1]-vv[1]))
                mymap[key]++
            }
        }
        for _, count := range mymap{
            answer += count * (count-1)
        }
    }
    return answer
}
