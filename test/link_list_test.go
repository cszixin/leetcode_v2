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
