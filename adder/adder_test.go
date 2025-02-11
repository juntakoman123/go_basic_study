package adder

import "testing"

func Test_addNumbers(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 + 1 = 2",
			args: args{
				x: 1,
				y: 1,
			},
			want: 2,
		},
		{
			name: "1 + 2 = 3",
			args: args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addNumbers(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("addNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
