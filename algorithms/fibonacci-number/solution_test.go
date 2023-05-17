package fibonacci_number

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{n: 2},
			1,
		},
		{
			"Example 2",
			args{n: 3},
			2,
		},
		{
			"Example 3",
			args{n: 4},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
