package _3096_minimumLevels

import "testing"

func Test_minimumLevels(t *testing.T) {
	type args struct {
		possible []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				possible: []int{1, 0, 1, 0},
			},
			want: 1,
		},
		{
			name: "t2",
			args: args{
				possible: []int{1, 1, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "t3",
			args: args{
				possible: []int{0, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLevels(tt.args.possible); got != tt.want {
				t.Errorf("minimumLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
