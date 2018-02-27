func sortList(head *ListNode) *ListNode {
	if head != nil {
		quicksort(head, nil)
	}
	return head
}

func quicksort(start *ListNode, end *ListNode) {

	value := start.Val

	cur := start
	prev := start // little

	for cur != end {
		if cur.Val < value {
			cur.Val, prev.Val = prev.Val, cur.Val
			if prev.Next != end {
				prev = prev.Next
			}
		}
		cur = cur.Next
	}
	if start != prev && start.Next != prev {
		quicksort(start, prev)
	}

	if prev.Val != value {
		quicksort(prev, end)
	} else {
		if prev.Next != end {
			quicksort(prev.Next, end)
		}
	}
}
