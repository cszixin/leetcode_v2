package myheap

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		// {"test1", args{[]int{5, 4, 3, 2, 1}}, []string{}},
		{"test2", args{[]int{10, 3, 8, 9, 4}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRelativeRanksv2(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks() = %v, want %v", got, tt.want)
			}
		})
	}
}
