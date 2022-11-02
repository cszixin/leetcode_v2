package sort

import "testing"

func TestMergerSortV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, -1, 3, 6, 8, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergerSortV2(tt.args.nums)
			print(12)
		})
	}
}
