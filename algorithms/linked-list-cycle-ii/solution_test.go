package linked_list_cycle_ii

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
	// Creating nodes for the first test case
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}

	// Linking nodes together to create a cyclic linked list
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // Creating a cycle

	// Creating nodes for the second test case
	n5 := &ListNode{Val: 1}
	n6 := &ListNode{Val: 2}

	// Linking nodes together to create a cyclic linked list
	n5.Next = n6
	n6.Next = n5 // Creating a cycle

	// Creating a node for the third test case
	n7 := &ListNode{Val: 1}

	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Test Case 1",
			head: n1,
			want: n2,
		},
		{
			name: "Test Case 2",
			head: n5,
			want: n5,
		},
		{
			name: "Test Case 3",
			head: n7,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.head); got != tt.want {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
