package knapsack

import "testing"

var (
	item11 = item{weight: 1, value: 1}
	item12 = item{weight: 1, value: 2}
	item21 = item{weight: 2, value: 1}
	item22 = item{weight: 2, value: 2}
	item24 = item{weight: 2, value: 4}
	item25 = item{weight: 2, value: 5}
	item38 = item{weight: 3, value: 8}
)

func TestGetMaxValueGivenMaxCapacity(t *testing.T) {
	type args struct {
		capacity int
		items    []item
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no capacity",
			args: args{capacity: 0, items: []item{item11}},
			want: 0,
		},
		{
			name: "1 capacity",
			args: args{capacity: 1, items: []item{item11, item12}},
			want: 2,
		},
		{
			name: "2 capacity",
			args: args{capacity: 2, items: []item{item11, item12, item21, item22}},
			want: 3,
		},
		{
			name: "4 capacity",
			args: args{capacity: 4, items: []item{item24, item25, item38}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxValueGivenMaxCapacity(tt.args.capacity, tt.args.items); got != tt.want {
				t.Errorf("GetMaxValueGivenMaxCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}
