//
//date: 2017-7-3
//desc: easy way
//

/*
1()(2)
1(1)
1(2)(3)
*/

func tree2str(t *TreeNode) string {
    if t == nil {
        return ""
    }
    
    str := fmt.Sprintf("%d",t.Val)  
    if t.Left != nil {
        left := tree2str(t.Left)
        str = str + "(" + left + ")"
        if t.Right != nil {
            right := tree2str(t.Right)
            str = str + "(" + right + ")"
        }
        return str
    }
    if t.Right != nil {
        right := tree2str(t.Right)
        str = str + "()" + "(" + right + ")"
    }   
    return str
}

