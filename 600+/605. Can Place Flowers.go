//
//date: 2017-6-12
//desc: func(num of 0) = (num - 1) / 2
//


func canPlaceFlowers(flowerbed []int, n int) bool {
    
    count := 0
    sum := 0
    
    if flowerbed[0] == 0 {
        count = 1
    }
    
    for _, v := range flowerbed {
        switch {
            case v == 0 :
                count++
            case count > 0 :
                sum += (count - 1) >> 1
                if sum >= n { return true}
                count = 0
        }
    }
    
    sum +=  count >> 1
       
    if sum >= n {
        return true
    } else {
        return false
    }
}
