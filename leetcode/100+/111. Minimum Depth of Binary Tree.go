//
//date: 2017-6-23
//desc: hello
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) int {
    switch {
        case left == nil && right == nil:
            return 0
        case left == nil && right != nil:
            return 1 + check(right.Left, right.Right)
        case left != nil && right == nil:
            return 1 + check(left.Left, left.Right)
        default:
            deepLeft  :=  1 + check(left.Left, left.Right)
            deepRight :=  1 + check(right.Left, right.Right)
            if deepLeft > deepRight {
                return deepRight
            }
            return deepLeft
    }
}
