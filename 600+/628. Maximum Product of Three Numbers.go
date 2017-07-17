//
//date: 2017-7-14
//desc: 1.这个程序本质也就是裂为了组n个数字，维持状态也就是也可以带n，最后的结果处理带n。
//

func min(nums *[]int, i int) {
    if (*nums)[i] > (*nums)[i + 1] {
        (*nums)[i], (*nums)[i + 1] = (*nums)[i + 1], (*nums)[i]
    } 
}

func sort3 (nums *[]int) {
    min(nums,0)
    min(nums,1)
    min(nums,0)
}

func sort(nums *[]int, next int) {
    if len(*nums) < 3 {
        *nums = append(*nums,next) 
        if len(*nums) == 3 {
            sort3(nums)
        }
        return
    }
    
    if next > (*nums)[0] {
        if next <= (*nums)[1] {
            (*nums)[0] = next
        }else if next <= (*nums)[2] {
            (*nums)[0], (*nums)[1] = (*nums)[1], next
        }else {
            (*nums)[0], (*nums)[1] = (*nums)[1], (*nums)[2]
            (*nums)[2] = next
        }
    }
    
}

func maximumProduct(nums []int) int {
    if len(nums) == 3 {
        return nums[0]*nums[1]*nums[2]
    }
    more1 := make([]int,0)
    less1 := make([]int,0)
    more := &more1
    less := &less1
    //var less *([]int) = nil
    
    for _, v := range nums {
        if v >= 0 {
            sort(more,v)
        }else{
            sort(less,-v)
        }
    }
    
    a := len(more1)
    b := len(less1)
    
    //fmt.Printf("%v \n",more1)
    //fmt.Printf("%v \n",less1)
    
    max := more1[a-1]
    c := 0
    d := 0
    
    if a > 2 {
        c = more1[0] * more1[1]
    }
    if b > 1 {
        d = less1[b-1]*less1[b-2]
    }
    
    if c > d {
        return c * max
    }else {
        return d * max
    }
}
