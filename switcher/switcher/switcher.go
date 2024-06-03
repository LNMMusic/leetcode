package switcher

type ListNode struct {
	Value int
	Next  *ListNode
}

// Swap
func Swap(head *ListNode) (swapped *ListNode) {
	// create a dummy node
	dummy := ListNode{Value: 0, Next: head}

	// ref node for iteration
	// - always 1 step behind the 2 followed nodes
	var ref_node *ListNode = &dummy
	for ref_node.Next != nil && ref_node.Next.Next != nil {
		// fetch nodes and re-wire
		start := ref_node.Next
		end := start.Next
		start.Next = end.Next
		end.Next = start

		// update chain
		ref_node.Next = end

		// update ref_node: always 2 nodes before
		ref_node = start
	}

	swapped = dummy.Next
	return
}
