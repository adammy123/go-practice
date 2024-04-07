package gd

import (
	"reflect"
	"testing"
)

const (
	L_1 = "L_1"
	L_2 = "L_2"
	L_3 = "L_3"
)

func TestPal(t *testing.T) {
	type args struct {
		header    string
		supported []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "no match",
			args: args{header: "L_1", supported: []string{L_2, L_3}},
			want: []string{},
		},
		{
			name: "L_1 match only",
			args: args{header: "L_2,L_1", supported: []string{L_1, L_3}},
			want: []string{L_1},
		},
		{
			name: "L_2 then L_3",
			args: args{header: "L_2,L_3", supported: []string{L_1, L_2, L_3}},
			want: []string{L_2, L_3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pal(tt.args.header, tt.args.supported); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pal() = %v, want %v", got, tt.want)
			}
		})
	}
}
