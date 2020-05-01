//
// date: 2017-7-10
// desc: hello
//

func findMode(root *TreeNode) []int {
    temp := make([]int,0)
    find(root,&temp)
    
    answer := make([]int, 0)
    count := make(map[int]int)
    max := 1
    
    for _, v := range temp {
        count[v]++
        if count[v] > max { max = count[v] }
    }
    
    for k, v := range count {
        if v == max { answer = append(answer,k) }
    }   
    return answer
}

func find(root *TreeNode, temp *[]int)  {
    if root == nil {
        return 
    } 
    find(root.Left, temp)
    *temp = append(*temp, root.Val)
    find(root.Right, temp)
}
