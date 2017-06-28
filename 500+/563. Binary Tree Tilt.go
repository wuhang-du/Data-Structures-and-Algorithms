//
//date: 2017-6-28
//desc: 以后这种题就直接找答案，就可以了。
//

func findTilt(root *TreeNode) int {
    answer := 0
    get(root, &answer)
    return answer
}
func get(root * TreeNode, sum *int) int {
    if root == nil {
        return 0
    }
    
    left := 0
    answer := 0
    right := 0
    
    left = get(root.Left, sum)
    right = get(root.Right, sum)
    switch {
        case left > right:
                answer = left - right
        case left < right:
                answer = right - left
        default:
                answer = 0
    }
    *sum += answer
    return root.Val + left + right
}
