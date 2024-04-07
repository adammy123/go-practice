package targetsum

import "testing"

func Test_changeSigns(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{nums: []int{1, 1, 1, 1, 1}, target: 3},
			want: 5,
		},
		{
			name: "test 2",
			args: args{nums: []int{1, 2, 1}, target: 2},
			want: 2,
		},
		{
			name: "test 3",
			args: args{nums: []int{1, 2, 7, 1, 5, 3}, target: 8},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeSigns(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("changeSigns() = %v, want %v", got, tt.want)
			}
		})
	}
}
