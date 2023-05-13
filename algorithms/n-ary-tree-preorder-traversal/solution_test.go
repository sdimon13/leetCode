package n_ary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example 1",
			args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 3, Children: []*Node{
							{Val: 5, Children: nil},
							{Val: 6, Children: nil},
						}},
						{Val: 2, Children: nil},
						{Val: 4, Children: nil},
					},
				},
			},
			[]int{1, 3, 5, 6, 2, 4},
		},
		{
			"Example 2",
			args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 2, Children: nil},
						{Val: 3, Children: []*Node{
							{Val: 6, Children: nil},
							{Val: 7, Children: []*Node{
								{Val: 11, Children: []*Node{
									{Val: 14, Children: nil},
								}},
							}},
						}},
						{Val: 4, Children: []*Node{
							{Val: 8, Children: []*Node{
								{Val: 12, Children: nil},
							}},
						}},
						{Val: 5, Children: []*Node{
							{Val: 9, Children: []*Node{
								{Val: 13, Children: nil},
							}},
							{Val: 10, Children: nil},
						}},
					},
				},
			},
			[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderIterative(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example 1",
			args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 3, Children: []*Node{
							{Val: 5, Children: nil},
							{Val: 6, Children: nil},
						}},
						{Val: 2, Children: nil},
						{Val: 4, Children: nil},
					},
				},
			},
			[]int{1, 3, 5, 6, 2, 4},
		},
		{
			"Example 2",
			args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{Val: 2, Children: nil},
						{Val: 3, Children: []*Node{
							{Val: 6, Children: nil},
							{Val: 7, Children: []*Node{
								{Val: 11, Children: []*Node{
									{Val: 14, Children: nil},
								}},
							}},
						}},
						{Val: 4, Children: []*Node{
							{Val: 8, Children: []*Node{
								{Val: 12, Children: nil},
							}},
						}},
						{Val: 5, Children: []*Node{
							{Val: 9, Children: []*Node{
								{Val: 13, Children: nil},
							}},
							{Val: 10, Children: nil},
						}},
					},
				},
			},
			[]int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderIterative(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
