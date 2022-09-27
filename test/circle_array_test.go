/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 03:40:01
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 03:51:26
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
