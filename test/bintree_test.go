/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:32:02
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-06 23:11:32
 */
package test

import (
	"leetcode/pkg/bintree"
	"testing"

	"github.com/stretchr/testify/assert"
)

var root *bintree.BTNode
var preorder []string
var inorder []string

func init() {
	preorder = []string{"A", "B", "D", "C", "E", "F", "G", "H", "I"}
	inorder = []string{"B", "C", "D", "A", "E", "G", "H", "F", "I"}
	root = bintree.CreateBinTree(preorder, inorder)

}

func TestPreOrderTree(t *testing.T) {
	res := bintree.PreOrder(root)
	assert.Equal(t, preorder, res, "")
}

func TestInOrderTree(t *testing.T) {
	res := bintree.InOrder(root)
	assert.Equal(t, inorder, res, "")
}

func TestPostOrderTree(t *testing.T) {
	res := bintree.Postrder(root)
	assert.Equal(t, []string{"C", "D", "B", "H", "G", "I", "F", "E", "A"}, res, "")
}

func TestLevelOrder(t *testing.T) {
	res := bintree.LevelOrder(root)
	assert.Equal(t, []string{"A"}, res, "")

}
