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

func TestReverseV1(t *testing.T) {
	input := []int{3, 6, 3, 5, 2, 1}
	head := link_list.CreateLinkList(input, false)
	new_head := link_list.ReverseV2(head)
	assert.Equal(t, []int{1, 2, 5, 3, 6, 3}, link_list.PrintLinkList(new_head))
}

func TestReverseV2(t *testing.T) {
	input := []int{3, 3, 1, 45, 6, 2, 34}
	head := link_list.CreateLinkList(input, false)
	head = link_list.ReverseV2(head)
	assert.Equal(t, []int{34, 2, 6, 45, 1, 3, 3}, link_list.PrintLinkList(head))
}

func TestReverseN(t *testing.T) {
	input := []int{3, 3, 1, 45, 6, 2, 34}
	head := link_list.CreateLinkList(input, false)
	head = link_list.ReverseN(head, 7)
	assert.Equal(t, []int{3, 3, 1, 45, 6, 2, 34}, link_list.PrintLinkList(head))

}

func TestReversev3(t *testing.T) {
	input := []int{4, 6, 1, 4, 78, 12, 5}
	expect := []int{5, 12, 78, 4, 1, 6, 4, 12}
	head := link_list.CreateLinkList(input, false)
	head = link_list.Reversev3(head)
	actual := link_list.PrintLinkList(head)
	assert.Equal(t, expect, actual, "")
}

func TestReverseBetween(t *testing.T) {
	input := []int{5, 12, 78, 4, 1, 6, 4, 12}
	head := link_list.CreateLinkList(input, false)
	a := link_list.GetNode(head, 2)
	b := link_list.GetNode(head, 4)
	new_head := link_list.ReverseBetween(a, b)
	expect := []int{12, 5}
	assert.Equal(t, expect, link_list.PrintLinkList(new_head), "")
}

func TestReverKGroup(t *testing.T) {
	input := []int{5, 12, 78, 4, 1, 6, 4, 12}
	k := 2
	expect := []int{12, 5, 4, 78, 6, 1, 12, 4}
	head := link_list.CreateLinkList(input, false)
	head = link_list.ReverKGroup(head, k)
	assert.Equal(t, expect, link_list.PrintLinkList(head), "")
}
