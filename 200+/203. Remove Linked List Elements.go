//
//date: 2017-6-11
//desc: 纸上得来终觉浅，觉知此事要躬行。
//


func removeElements(head *ListNode, val int) *ListNode {
    point := &head
    for *point != nil {
        if (**point).Val != val {
            point = &((**point).Next)
        } else {
            *point = (**point).Next
        }
    } 
    return head
}
