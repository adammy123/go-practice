package firstunique

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{A: []int{4, 10, 5, 4, 2, 10}},
			want: 5,
		},
		{
			name: "test 2",
			args: args{A: []int{1, 4, 3, 3, 1, 2}},
			want: 4,
		},
		{
			name: "test 3",
			args: args{A: []int{1, 2, 2, 1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
