package noderemoval

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// count full length of listnode
	var cc int
	var ref *ListNode = head
	for ref != nil {
		cc++
		ref = ref.Next
	}

	// get element to pop
	nn := cc - n + 1

	// pop out element
	dummy := &ListNode{Val: 0, Next: head}
	cc = 0
	ref = dummy // ref starts from current and keeping an eye on the next as the one that must be popped
	for ref != nil {
		cc++
		// we check if the cc is the n° of element equally to the desire to be popped
		if cc == nn {
			// re-wire: set the followed nodes of the pop to the current ref
			nextNode := ref.Next
			ref.Next = nextNode.Next
			// disconnect: to avoid this pop node keep tracking of nodes that should not
			nextNode.Next = nil
			break
		}

		// update ref
		ref = ref.Next
	}

	return dummy.Next
}
