package bulls_and_cows

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Example 1",
			args{
				secret: "1807",
				guess:  "7810",
			},
			"1A3B",
		},
		{
			"Example 2",
			args{
				secret: "1123",
				guess:  "0111",
			},
			"1A1B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
