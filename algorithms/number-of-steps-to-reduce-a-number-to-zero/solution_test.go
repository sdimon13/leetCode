package number_of_steps_to_reduce_a_number_to_zero

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				14,
			},
			6,
		},
		{
			"Example 2",
			args{
				8,
			},
			4,
		},
		{
			"Example 3",
			args{
				123,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberOfStepsV2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				14,
			},
			6,
		},
		{
			"Example 2",
			args{
				8,
			},
			4,
		},
		{
			"Example 3",
			args{
				123,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfStepsV2(tt.args.num); got != tt.want {
				t.Errorf("numberOfStepsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
