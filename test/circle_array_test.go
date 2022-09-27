/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 03:40:01
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 05:52:10
 */
package test

import (
	"leetcode/pkg/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 0, 1, 1, 2}
	expect := 3
	actual := array.RemoveDuplicates(input)
	assert.Equal(t, expect, actual, "")

	expect_array := []int{0, 1, 2, 2, 2}
	assert.Equal(t, expect_array, input, "")

}

func TestRemoveElement(t *testing.T) {
	type args struct {
		array []int
		val   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 3, 1, 1, 3}, 3}, 3},
		{"test2", args{[]int{0, 0, 1, 1, 1, 3}, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.RemoveElement(tt.args.array, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
