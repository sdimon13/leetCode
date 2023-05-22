package top_k_frequent_words

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Example 1",
			args{
				words: []string{
					"i", "love", "leetcode", "i", "love", "coding",
				},
				k: 2,
			},
			[]string{"i", "love"},
		},
		{
			"Example 2",
			args{
				words: []string{
					"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is",
				},
				k: 4,
			},
			[]string{"the", "is", "sunny", "day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
