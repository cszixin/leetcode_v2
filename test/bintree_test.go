/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:32:02
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-07 19:55:03
 */
package test

import (
	"leetcode/pkg/bintree"
	"reflect"
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
	assert.Equal(t, []string{"A", "B", "E", "D", "F", "C", "G", "I", "H"}, res, "")

}

func TestGetNodeLevel(t *testing.T) {
	type args struct {
		root  *bintree.BTNode
		value string
		level int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{root, "F", 1}, 3},
		{"test2", args{root, "O", 1}, -1},
		{"test3", args{root, "H", 1}, 5},
		{"test4", args{root, "A", 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetNodeLevel(tt.args.root, tt.args.value, tt.args.level); got != tt.want {
				t.Errorf("GetNodeLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNode(t *testing.T) {
	type args struct {
		root  *bintree.BTNode
		value string
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{root, "B"}, root.Left},
		{"test2", args{root, "A"}, root},
		{"test3", args{root, "C"}, root.Left.Right.Left},
		{"test4", args{root, "P"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.FindNode(tt.args.root, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNode() = %v, want %v value = %v", got, tt.want, got.Data)
			}
		})
	}
}

func TestGetBintreeDepth(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{root}, 5},
		{"test2", args{root.Left}, 3},
		{"test3", args{root.Right.Right}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetBintreeDepth(tt.args.root); got != tt.want {
				t.Errorf("GetBintreeDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetKthNode(t *testing.T) {
	type args struct {
		root *bintree.BTNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{root, 2}, root.Left},
		{"test2", args{root, 1}, root},
		{"test3", args{root, 10}, nil},
		{"test4", args{root, 9}, root.Right.Right.Right},
		{"test5", args{root, 8}, root.Right.Right.Left.Right},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetKthNode(tt.args.root, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKthNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTreeNodeNum(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{root}, 9},
		{"test2", args{root.Left}, 3},
		{"test3", args{root.Left.Right.Left}, 1},
		{"test4", args{root.Left.Right.Left.Left}, 0},
		{"test5", args{root.Right}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetTreeNodeNum(tt.args.root); got != tt.want {
				t.Errorf("GetTreeNodeNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
