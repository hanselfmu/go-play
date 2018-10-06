package algorithm

// ReverseNodePairs swaps every two adjacent nodes and return its head.
// LeetCode #24
func ReverseNodePairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	cur := head
	head = cur.Next
	var prev *ListNode
	for cur != nil && cur.Next != nil {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = cur

		if prev != nil {
			prev.Next = nxt
		}
		prev = cur

		cur = cur.Next

	}

	return head
}

// ReverseKNodes reverses K adjacent nodes and return its head.
// LeetCode #25
func ReverseKNodes(head *ListNode, k int) *ListNode {
	size := 0
	n := head
	for n != nil {
		size++
		n = n.Next
	}
	if size < 2 {
		return head
	}
	idx := 0
	limit := size - size%k
	cur := head.Next
	var prev *ListNode
	// todo
	for idx < limit {
		nxt := cur.Next
		prev.Next = nxt

		idx++
	}

	return head
}
