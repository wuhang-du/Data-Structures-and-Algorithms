//
//date: 2017-6-20
//desc: 
//

func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    check(root, &sum)
    return sum
}
 
func check(root *TreeNode, sum *int) {
    if root == nil { return }
    if root.Left != nil {
        if root.Left.Left == nil && root.Left.Right == nil {
            *sum += root.Left.Val
        } else {
            check(root.Left, sum)
        }
    }
    
    if root.Right != nil {
        check(root.Right, sum)
    }
}
