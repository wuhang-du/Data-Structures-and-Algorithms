/**
 * Definition for a binary tree node.
  * type TreeNode struct {
	   *     Val int
	    *     Left *TreeNode
		 *     Right *TreeNode
		  * }
*/
func rightSideView(root *TreeNode) []int {
	answer := []int{}
	view(root, &answer, 0)
	return answer
}

func view(root *TreeNode, answer *[]int, cur int) {
	if root == nil {
		return
	}
	if cur == len(*answer) {
		*answer = append(*answer, root.Val)
	}

	view(root.Right, answer, cur+1)
	view(root.Left, answer, cur+1)

}
