package merge_strings_alternately

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Example 1",
			args{
				word1: "abc",
				word2: "pqr",
			},
			"apbqcr",
		},
		{
			"Example 2",
			args{
				word1: "ab",
				word2: "pqrs",
			},
			"apbqrs",
		},
		{
			"Example 3",
			args{
				word1: "abcd",
				word2: "pq",
			},
			"apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
