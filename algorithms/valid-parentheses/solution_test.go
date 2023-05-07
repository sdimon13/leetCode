package valid_parentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{s: "()"},
			true,
		},
		{
			"Example 2",
			args{s: "()[]{}"},
			true,
		},
		{
			"Example 3",
			args{s: "(]"},
			false,
		},
		{
			"Example 4",
			args{s: "([{}])"},
			true,
		},
		{
			"Example 5",
			args{s: "([{})]"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
