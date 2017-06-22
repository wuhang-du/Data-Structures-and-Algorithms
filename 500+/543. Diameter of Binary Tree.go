//
//date: 2017-6-22
//desc: 先根遍历，要多算数据
//		后根遍历，之前的数据后面可以使用
//

func diameterOfBinaryTree(root *TreeNode) int {
    sum := 0
    _ = lastRoot(root, &sum)
    return sum
}

func lastRoot(root *TreeNode, sum *int) int {
    if root == nil {
        return 0
    }

    left := lastRoot(root.Left, sum)
    right := lastRoot(root.Right, sum)
    
    deepth := left
    if left < right {
        deepth = right
    }
    deepth++
    if left + right > *sum {
        *sum = left + right
    }
    return deepth
}



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    
    sum := 0
    root_first(root, &sum)
    return sum
}

func root_first(root *TreeNode, sum *int) {
    
    if root == nil {
        return
    }
    
    a := root_check(root)
    
    if a > *sum {
        *sum = a
    }
    if root.Left != nil {
        root_first(root.Left, sum)
    }
    if root.Right != nil {
        root_first(root.Right,sum)
    }
}



func root_check (root * TreeNode) int {
    left := 0
    if root.Left != nil {
        left =  1 + max_num(root.Left.Left, root.Left.Right)
    }
    
    right := 0
    if root.Right != nil {
        right =  1 + max_num(root.Right.Left, root.Right.Right)
    }
    return left + right
}


func max_num(left *TreeNode, right *TreeNode) int {
    
    if left == nil && right == nil {
        return 0
    }
    
    left_sum := 1
    if left != nil {
        left_sum += max_num(left.Left,left.Right)
    }
    right_sum := 1
    if right != nil {
        right_sum += max_num(right.Left,right.Right)
    }
    
    
    if left_sum < right_sum {
        return right_sum
    }
    return left_sum
}
