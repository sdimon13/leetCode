package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		temp := head.Next
		head.Next = reverse
		reverse = head
		head = temp
	}

	return reverse
}

func reverseListV2(head *ListNode) (reverse *ListNode) {
	for head != nil {
		head.Next, reverse, head = reverse, head, head.Next
	}

	return
}
