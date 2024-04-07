package tennistournament

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		P int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{P: 5, C: 3},
			want: 2,
		},
		{
			name: "test 2",
			args: args{P: 10, C: 3},
			want: 3,
		},
		{
			name: "test 3",
			args: args{P: 11, C: 10},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.P, tt.args.C); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
