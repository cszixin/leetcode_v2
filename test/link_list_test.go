package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/pkg/link_list"
)

func TestLinkList(t *testing.T) {
	input := []int{5, 2, 1, 6, 4}
	node := link_list.CreateLinkList(input, false)
	res := link_list.PrintLinkList(node)
	assert.Equal(t, input, res, "")
}

func TestMiddleNode(t *testing.T) {
	input := []int{5, 2, 1, 6, 4, 9}
	head := link_list.CreateLinkList(input, false)
	output := link_list.MiddleLinkList(head)
	expect := 6
	assert.Equal(t, expect, output, "should be the same")
}

func TestCycleLink(t *testing.T) {
	input := []int{5, 2, 1, 6, 4, 9}
	head := link_list.CreateLinkList(input, true)
	output := link_list.HasCycle(head)
	expect := true
	assert.Equal(t, expect, output, "")

}

func TestCreateIntersection(t *testing.T) {
	h1, h2 := link_list.CreateIntersection()
	output1 := link_list.PrintLinkList(h1)
	expect1 := []int{3, 6, 3, 1, 12, 13, 14}
	assert.Equal(t, expect1, output1, "")

	output2 := link_list.PrintLinkList(h2)
	expect2 := []int{6, 7, 8, 10, 12, 13, 14}

	assert.Equal(t, expect2, output2, "")
}

func TestIfIntersection(t *testing.T) {
	h1, h2 := link_list.CreateIntersection()
	value, output := link_list.IfIntersection(h1, h2)
	assert.Equal(t, true, output, "")
	assert.Equal(t, 12, value, "")

}
