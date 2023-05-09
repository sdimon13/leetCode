package isomorphic_strings

import "testing"

func Test_isIsomorphic(t *testing.T) {
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
				s: "egg",
				t: "add",
			},
			true,
		},
		{
			"Example 2",
			args{
				s: "foo",
				t: "bar",
			},
			false,
		},
		{
			"Example 3",
			args{
				s: "paper",
				t: "title",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIsomorphicV2(t *testing.T) {
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
				s: "egg",
				t: "add",
			},
			true,
		},
		{
			"Example 2",
			args{
				s: "foo",
				t: "bar",
			},
			false,
		},
		{
			"Example 3",
			args{
				s: "paper",
				t: "title",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphicV2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphicV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
