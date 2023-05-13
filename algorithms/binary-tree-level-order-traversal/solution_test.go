package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Example 1",
			args{root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}},
			[][]int{
				{
					3,
				},
				{
					9, 20,
				},
				{
					15, 7,
				},
			},
		},
		{
			"Example 2",
			args{&TreeNode{Val: 1}},
			[][]int{
				{1},
			},
		},
		{
			"Example 3",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
