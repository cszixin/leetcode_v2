/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-19 09:38:48
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 11:22:16
 */
package sort

import (
	"reflect"
	"testing"
)

func TestQuickSortV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{-1, 2, -5, 6, 9, 3, 12, 0, 3}}, []int{-5, -1, 0, 2, 3, 3, 6, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargest(t *testing.T) {
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
		{"test1", args{[]int{1, 2, -4, 4, -3, 9}, 2}, 4},
		{"test2", args{[]int{1, 2, -4, 4, -3, 9}, 1}, 9},
		{"test3", args{[]int{1, 2, -4, 4, -3, 9}, 7}, -1},
		{"test4", args{[]int{1, 2, -4, 4, -3, 9}, 6}, -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
