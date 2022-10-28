package test

import (
	"leetcode/pkg/myheap"
	"reflect"
	"testing"
)

func TestSortByHeap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{4, 6, 1, 2, 7}}, []int{7, 6, 4, 2, 1}},
		{"test2", args{[]int{0, 1, 1, 1}}, []int{1, 1, 1, 0}},
		{"test3", args{[]int{1, 1, 1, 1}}, []int{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myheap.SortByHeap(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindKth(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 8, 2, 5, 4}, 3}, 4},
		{"test2", args{[]int{-1, 0, -3, 2, 5, 10}, 5}, 5},
		{"test2", args{[]int{-1, 0, -3, 2, 5, 10}, 6}, 10},
		{"test2", args{[]int{-1, 0, -3, 2, 5, 10}, 1}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myheap.FindKth(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
