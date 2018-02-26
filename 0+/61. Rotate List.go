/**
 * Definition for singly-linked list.
  * type ListNode struct {
	   *     Val int
	    *     Next *ListNode
		 * }
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return head
	}

	cur := head
	count := 1

	for cur.Next != nil {
		count++
		cur = cur.Next
	}

	k = k % count

	cur.Next = head
	limit := count - k
	cur = head

	for limit > 1 {
		cur = cur.Next
		limit--
	}

	head = cur.Next
	cur.Next = nil
	return head
}
