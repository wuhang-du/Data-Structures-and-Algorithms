// 首先使用的是 尾递归
// 之后改成了 循环，但是 视角是第二个节点
// 之后改成了 第一个节点 的视角。

func swapPairs(head *ListNode) *ListNode {
	tmp := ListNode{}
	tmp.Next = head
	now := &tmp
	var empty *ListNode

	for now.Next != empty && now.Next.Next != empty {

		first := now.Next
		second := now.Next.Next

		now.Next = second
		first.Next = second.Next
		second.Next = first

		now = now.Next.Next
	}

	return tmp.Next
}
