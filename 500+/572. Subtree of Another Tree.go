//
//date： 2017-6-15
//desc:  两层递归，所以有了重复的代码。
//

func isSubtree(s *TreeNode, t *TreeNode) bool {
    switch {
        case s == nil && t == nil:
            return true
        case s != nil && t == nil:
            return false
        case s == nil && t != nil:
            return false
    }
    
    if Check(s, t) { return true}
    
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func Check(s *TreeNode, t *TreeNode) bool {
     switch {
        case s == nil && t == nil:
            return true
        case s != nil && t == nil:
            return false
        case s == nil && t != nil:
            return false
    }
    
    if s.Val == t.Val {
        return Check(s.Left, t.Left) && Check(s.Right, t.Right)
    }   
    return false
}
