package fizz_buzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				n: 3,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
			},
		},
		{
			name: "Example 2",
			args: args{
				n: 5,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
			},
		},
		{
			name: "Example 3",
			args: args{
				n: 15,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fizzBuzz(tt.args.n))
		})
	}
}

func Test_fizzBuzzV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				n: 3,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
			},
		},
		{
			name: "Example 2",
			args: args{
				n: 5,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
			},
		},
		{
			name: "Example 3",
			args: args{
				n: 15,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fizzBuzzV2(tt.args.n))
		})
	}
}
