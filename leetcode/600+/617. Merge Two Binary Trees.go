//
//date: 2017-6-26
//desc: gradually optimize
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    //将t1 merge 到 t2 上去
    switch {
        case t1 == nil :
        case t2 == nil :
            t2 = t1
        case t2 != nil :
            t2.Val += t1.Val
            t2.Left = mergeTrees(t1.Left, t2.Left)
            t2.Right = mergeTrees(t1.Right, t2.Right)
    }
    return t2
}
