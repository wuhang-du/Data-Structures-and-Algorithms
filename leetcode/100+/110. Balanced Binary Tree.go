//
//date: 2017-6-26
//desc: 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
// 后序遍历
func isBalanced(root *TreeNode) bool {
   if Getdeepth(root) == -1 {
       return false
   }
   return true
}


func Getdeepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    dl := Getdeepth(root.Left)
    dr := Getdeepth(root.Right)
    
    if dl == -1 || dr == -1 {
        return -1
    }
    
    if dl < dr {
        dl, dr = dr, dl
    }
    
    if dl - dr > 1 {
        return -1
    }
    dl++
    return dl
}



