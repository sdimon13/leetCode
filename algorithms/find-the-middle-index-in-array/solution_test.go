package find_the_middle_index_in_array

import "testing"

func Test_findMiddleIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{nums: []int{
				2, 3, -1, 8, 4,
			}},
			3,
		},
		{
			"Example 2",
			args{nums: []int{
				1, -1, 4,
			}},
			2,
		},
		{
			"Example 3",
			args{nums: []int{
				2, 5,
			}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
