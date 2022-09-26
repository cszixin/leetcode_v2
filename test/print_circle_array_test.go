/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-26 06:24:58
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-26 06:24:58
 */
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/pkg/array"
)

func TestPrintCircleArray(t *testing.T) {
	input := []int{2, 1, 3, 5, 6}
	expect := []int{2, 1, 3, 5, 6, 2, 1, 3, 5, 6}
	output := array.PrintCircleArray(input)
	assert.Equal(t, expect, output, "")
}
