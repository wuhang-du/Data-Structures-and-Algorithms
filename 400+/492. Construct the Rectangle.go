//
//date: 2017-6-19
//desc: easy
//

func constructRectangle(area int) []int {
    a := int(math.Sqrt(float64(area)))
    for area % a != 0 {
        a--
    }
    return []int{area/a,a}
}
