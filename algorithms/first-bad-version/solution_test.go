package first_bad_version

import "testing"

func TestVersionControl_firstBadVersion(t *testing.T) {
	type fields struct {
		badVersion int
	}
	type args struct {
		n            int
		isBadVersion func(int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Example 1",
			fields: fields{
				badVersion: 4,
			},
			args: args{
				n: 5,
			},
			want: 4,
		},
		{
			name: "Example 2",
			fields: fields{
				badVersion: 1,
			},
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewVersionControl(tt.fields.badVersion)
			isBadVersion := func(version int) bool {
				return version >= v.badVersion
			}
			if got := v.firstBadVersion(tt.args.n, isBadVersion); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
