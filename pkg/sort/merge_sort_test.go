package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 5, 3, -1, 9, 2, 10}}, []int{-1, 1, 2, 3, 5, 9, 10}},
		{"test2", args{[]int{1, 1, 1, 1}}, []int{1, 1, 1, 1}},
		{"test3", args{[]int{1, 1, 1, 1, 0}}, []int{0, 1, 1, 1, 1}},
		{"test4", args{[]int{0, 1, 1, 1, 1, 0}}, []int{0, 0, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallMergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
