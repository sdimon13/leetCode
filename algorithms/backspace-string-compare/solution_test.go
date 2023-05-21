package backspace_string_compare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{
				s: "ab#c",
				t: "ad#c",
			},
			true,
		},
		{
			"Example 2",
			args{
				s: "ab##",
				t: "c#d#",
			},
			true,
		},
		{
			"Example 3",
			args{
				s: "a#c",
				t: "b",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
