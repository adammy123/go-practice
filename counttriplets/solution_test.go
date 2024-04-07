package solution

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int64
		r   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test 1",
			args: args{arr: []int64{1, 2, 2, 4}, r: 2},
			want: 2,
		},
		{
			name: "test 2",
			args: args{arr: []int64{1, 3, 9, 9, 27, 81}, r: 3},
			want: 6,
		},
		{
			name: "test 5",
			args: args{arr: []int64{1, 5, 5, 25, 125}, r: 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr, tt.args.r); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
