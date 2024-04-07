package solution

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		S string
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "code",
			args: args{S: "cdeo", A: []int{3, 2, 0 ,1}},
			want: "code",
		},
		{
			name: "centipede",
			args: args{S: "cdeenetpi", A: []int{5, 2, 0, 1, 6, 4, 8, 3, 7}},
			want: "centipede",
		},
		{
			name: "bat",
			args: args{S: "bytdag", A: []int{4, 3, 0, 1, 2, 5}},
			want: "bat",
		},
		{
			name: "a",
			args: args{S: "abcde", A: []int{0, 1, 2, 3, 4}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S, tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
