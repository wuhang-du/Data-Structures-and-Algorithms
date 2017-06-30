//
//date: 2017-7-1
//desc: "" 切片的递归
//


import (
    "fmt"
    "strconv"
)

func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    answer := make([]string,0)
    temp := strconv.Itoa(root.Val)
    find(root, &answer, temp)
    return answer
}

func find(root *TreeNode,answer *[]string, temp string) {
    
    if root.Left == nil && root.Right == nil {
        *answer = append(*answer,temp)
        return
    }
    if root.Left != nil {
        temp1 := fmt.Sprintf("%s->%d",temp,(root.Left.Val))
        find(root.Left, answer, temp1)
    }
    if root.Right != nil {
        temp2 := fmt.Sprintf("%s->%d",temp,(root.Right.Val))
        find(root.Right, answer, temp2)
    }   
}
