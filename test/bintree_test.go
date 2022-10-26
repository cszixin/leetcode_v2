/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:32:02
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 15:50:06
 */
package test

import (
	"leetcode/pkg/bintree"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var root *bintree.BTNode
var rootA *bintree.BTNode
var preorder []string
var inorder []string

func init() {
	preorder = []string{"A", "B", "D", "C", "E", "F", "G", "H", "I"}
	inorder = []string{"B", "C", "D", "A", "E", "G", "H", "F", "I"}
	root = bintree.CreateBinTree(preorder, inorder)
	preorderA := []string{"0"}
	inorderA := []string{"0"}
	rootA = bintree.CreateBinTree(preorderA, inorderA)

}

func TestPreOrderTree(t *testing.T) {
	res := bintree.PreOrder(root)
	assert.Equal(t, preorder, res, "")
}

func TestPreOrderNonrecursion(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test1", args{root}, preorder},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.PreOrderNonrecursion(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrderNonrecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInOrderTree(t *testing.T) {
	res := bintree.InOrder(root)
	assert.Equal(t, inorder, res, "")
}

func TestInOrderNonrecursion(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test1", args{root}, inorder},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.InOrderNonrecursion(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderNonrecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostOrderTree(t *testing.T) {
	res := bintree.PostOrder(root)
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
			if got := bintree.GetBintreeDepthV2(tt.args.root); got != tt.want {
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

func TestGetPreNode(t *testing.T) {
	type args struct {
		root *bintree.BTNode
		p    *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{root, root.Right}, root.Left.Right.Left},
		{"test2", args{root, root}, nil},
		{"test3", args{root, root.Right.Right.Right}, root.Right.Right.Left.Right},
		{"test4", args{root, root.Right.Right.Left}, root.Right.Right},
		{"test5", args{root, root.Left}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetPreNodev2(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPreNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertBinTree(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{root}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.RevertBinTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverBinTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostOrderNonrecursion(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test1", args{root}, []string{"C", "D", "B", "H", "G", "I", "F", "E", "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.PostOrderNonrecursion(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrderNonrecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLeftRightNode(t *testing.T) {
	type args struct {
		root *bintree.BTNode
		p    *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{root, root}, []int{3, 5}},
		{"test2", args{root, root.Left}, []int{0, 2}},
		{"test3", args{root, root.Left.Right}, []int{1, 0}},
		{"test4", args{root, root.Right}, []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetLeftRightNode(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLeftRightNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetdiameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{root}, 7},
		{"test2", args{root.Left}, 2},
		{"test3", args{root.Right}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetdiameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("GetdiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderV2(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"test1", args{root}, [][]string{{"A"}, {"B", "E"}, {"D", "F"}, {"C", "G", "I"}, {"H"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.LevelOrderV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertBinTreeV2(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{root}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bintree.RevertBinTreeV2(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevertBinTreeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecur(t *testing.T) {
	type args struct {
		A *bintree.BTNode
		B *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{root, rootA}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.Recur(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("Recur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *bintree.BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 2, 3}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.ConstructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllRoute(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"test1", args{root}, [][]string{{}}},
		{"test2", args{root.Right.Right}, [][]string{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.GetAllRoute(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *bintree.BTNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{rootA}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bintree.IsValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
