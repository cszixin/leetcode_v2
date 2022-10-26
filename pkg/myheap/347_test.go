package myheap

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"test2", args{[]int{1}, 1}, []int{1}},
		{"test3", args{[]int{-1, -1}, 1}, []int{-1}},
		{"test4", args{[]int{1, 2}, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentV2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
