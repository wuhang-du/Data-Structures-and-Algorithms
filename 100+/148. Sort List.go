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

func sortList(head *ListNode) *ListNode {
	return split(head)
}

func split(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil

	// head and fast respent two different list
	head = split(head)
	fast = split(fast)

	return merge(head, fast)
}

func merge(line1 *ListNode, line2 *ListNode) *ListNode {

	var l ListNode
	cur := &l

	for line1 != nil && line2 != nil {
		if line1.Val < line2.Val {
			cur.Next = line1
			line1 = line1.Next
		} else {
			cur.Next = line2
			line2 = line2.Next
		}
		cur = cur.Next
	}

	if line1 == nil {
		cur.Next = line2
	}

	if line2 == nil {
		cur.Next = line1
	}

	return l.Next
}
