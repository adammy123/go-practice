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
			args: args{A: []int{10, 1, 3, 1, 2, 2, 1, 0, 4}},
			want: 3,
		},
		{
			name: "test 2",
			args: args{A: []int{5, 3, 1, 3, 2, 3}},
			want: 1,
		},
		{
			name: "test 3",
			args: args{A: []int{9, 9, 9, 9, 9}},
			want: 2,
		},
		{
			name: "test 4",
			args: args{A: []int{1, 5, 2, 4, 3, 3}},
			want: 3,
		},
		{
			name: "test 5",
			args: args{A: makeA(100)},
			want: 50,
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

func makeA(N int) []int {
	randomNumber := 1
	A := make([]int, N)
	for i, _ := range A {
		A[i] = randomNumber
	}
	return A
}
