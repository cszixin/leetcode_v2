/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-28 17:32:36
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-29 14:02:37
 */
package test

import (
	"leetcode/pkg/search"
	"testing"
)

func TestBinsearch(t *testing.T) {
	type args struct {
		array []int
		val   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6, 7}, 3}, 1},
		{"test2", args{[]int{1, 3, 5, 6, 7}, 4}, -1},
		{"test3", args{[]int{1, 1, 5, 6, 7}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.BinSearch(tt.args.array, tt.args.val); got != tt.want {
				t.Errorf("Binsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearchLeftBound(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{0, 0, 0, 2, 2, 2, 2}, 0}, 0},
		{"test2", args{[]int{0, 0, 1, 2, 2, 2, 2}, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.BinSearchLeftBound(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinSearchLeftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearch33(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test2", args{[]int{4, 5, 6, 7, 8, 9, 12}, 12}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.BinSearch33(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinSearch33() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinSearchRightBound(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 1, 2}, 3}, -1},
		{"test2", args{[]int{1, 1, 2}, 1}, 1},

		{"test3", args{[]int{1, 1, 2, 2, 2, 4, 4}, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.BinSearchRightBound(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinSearchRightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
