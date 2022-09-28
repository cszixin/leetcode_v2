/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-28 01:22:52
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-28 01:22:53
 */

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/pkg/common"
)

func TestNextGreaterElement(t *testing.T) {
	input := []int{5, 3, 6, 4, 5, 8, 7}
	expect := []int{6, 6, 8, 5, 8, -1, -1}
	output := common.NextGreaterElement(input)
	assert.Equal(t, expect, output, "should be the same.")
}

func TestDailyTemperatures(t *testing.T) {
	input := []int{2, 3, 6, 4, 5, 8, 7}
	expect := []int{1, 1, 3, 1, 1, 0, 0}
	output := common.DailyTemperatures(input)
	assert.Equal(t, expect, output, "should be the same")
}
