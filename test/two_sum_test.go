/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 07:02:13
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 07:14:26
 */
package test

import (
	"leetcode/pkg/search"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"test2", args{[]int{2, 7, 11, 15}, 7}, []int{-1, -1}},
		{"test3", args{[]int{2, 7, 11, 15}, 26}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.TwoSum(tt.args.array, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
