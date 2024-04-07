package solution

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
			args: args{A: []int{1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2}},
			want: 2,
		},
		{
			name: "test 2",
			args: args{A: []int{3, 5}},
			want: 0,
		},
		{
			name: "test 3",
			args: args{A: []int{3, 1, 2}},
			want: 1,
		},
		{
			name: "test 4",
			args: args{A: []int{6, 5, 4, 1, 5}},
			want: 4,
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
