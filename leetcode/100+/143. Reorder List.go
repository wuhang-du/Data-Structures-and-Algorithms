/**
 * Definition for singly-linked list.
  * type ListNode struct {
	   *     Val int
	    *     Next *ListNode
		 * }
*/
func reorderList(head *ListNode) {
	var empty *ListNode
	if head == empty || head.Next == empty {
		return
	}

	mid := head
	end := head

	for end != empty && end.Next != empty {
		mid = mid.Next
		end = end.Next.Next
	}
	// get 1-2-3-4 ==》 1-2- 3 - null
	//                    4 /

	tmp := mid
	before := empty
	for tmp != empty {
		n := tmp.Next
		tmp.Next = before
		before = tmp
		tmp = n
	}
	// now merge two lists

	cur := head
	cen := before

	for cen != mid {
		t := cur.Next
		m := cen.Next
		cur.Next = cen
		cen.Next = t
		cur = t
		cen = m
	}
	return
}
