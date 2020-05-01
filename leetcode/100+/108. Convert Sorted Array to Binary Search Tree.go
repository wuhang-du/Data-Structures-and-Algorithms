//
//date: 2017-7-3
//desc: come on 猜测，尝试，验证。
//

func sortedArrayToBST(nums []int) *TreeNode {
    length := len(nums)
    switch length {
        case 0:
            return nil
        case 1:
            return &TreeNode{nums[0], nil, nil}
        case 2:
            return &TreeNode{nums[1],&TreeNode{nums[0],nil,nil},nil}
        case 3:
            return &TreeNode{nums[1],&TreeNode{nums[0],nil,nil},&TreeNode{nums[2],nil,nil}}
        default:
            return &TreeNode{nums[length/2], sortedArrayToBST(nums[:length/2]), sortedArrayToBST(nums[length/2 + 1:])}                }
}
