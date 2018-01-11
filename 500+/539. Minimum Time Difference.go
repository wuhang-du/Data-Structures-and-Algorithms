// S: o(N)
// T: o(N)
func findMinDifference(timePoints []string) int {  
    table := make([]int,1440,1440)
    
    minleft := 1440
    minright := 1440
    
    for _, v :=range timePoints {
        t := strings.Split(v, ":")
        hour, _ := strconv.Atoi(t[0])
        minute ,_ := strconv.Atoi(t[1])
        time := hour * 60 + minute
        
        switch table[time] {
            case 1:
            return 0
            default:
            table[time] = 1
        }
        
        if 1440 - time < minright {
            minright = 1440 - time
        }
        
        if time - 0 < minleft {
            minleft = time - 0
        }         
    }
    
    min := minright + minleft
    prev := minleft
    
    for i := minleft + 1; i <= 1440 - minright; i++ {
        if table[i] > 0 {
            tmp := i - prev;
            prev = i
            if tmp < min {
                min = tmp
            }
        }      
    }
    
    return min
}
