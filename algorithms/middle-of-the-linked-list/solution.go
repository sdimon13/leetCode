package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var listArray []*ListNode

	for head != nil {
		listArray = append(listArray, head)
		head = head.Next
	}

	return listArray[len(listArray)/2]
}

func middleNodeV2(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
