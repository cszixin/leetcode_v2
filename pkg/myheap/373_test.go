package myheap

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 7, 11}, []int{2, 4, 6}, 3}, [][]int{[]int{1, 2}, []int{1, 4}, []int{1, 6}}},
		{"test2", args{[]int{1, 1, 2}, []int{1, 2, 3}, 2}, [][]int{[]int{1, 1}, []int{1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
