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
