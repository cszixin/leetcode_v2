/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 12:42:16
 */
package bintree

import (
	"fmt"
	"math"
	"strconv"
)

func PreOrder(root *BTNode) []string {
	res := []string{}
	var dfs func(*BTNode)
	dfs = func(root *BTNode) {
		if root != nil {
			res = append(res, root.Data)
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}

func PreOrderNonrecursion(root *BTNode) []string {
	res := make([]string, 0)
	stack := make([]*BTNode, 0)
	p := root
	//根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		p = stack[len(stack)-1]
		res = append(res, p.Data)
		// 出栈
		stack = stack[:len(stack)-1]
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return res
}

func InOrder(root *BTNode) []string {
	res := []string{}
	var dfs func(*BTNode)
	dfs = func(root *BTNode) {
		if root != nil {
			dfs(root.Left)
			res = append(res, root.Data)
			dfs(root.Right)
		}
	}
	dfs(root)
	return res
}

func InOrderNonrecursion(root *BTNode) []string {
	res := make([]string, 0)
	stack := make([]*BTNode, 0)
	p := root
	for len(stack) > 0 || p != nil {
		//到达最左
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			res = append(res, p.Data)
			stack = stack[:len(stack)-1]
			p = p.Right

		}

	}
	return res
}

func PostOrder(root *BTNode) []string {
	res := []string{}
	var dfs func(*BTNode)
	dfs = func(root *BTNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			res = append(res, root.Data)
		}
	}
	dfs(root)
	return res
}

func PostOrderNonrecursion(root *BTNode) []string {
	res := make([]string, 0)
	stack := make([]*BTNode, 0)
	p := root
	stack = append(stack, p)
	for len(stack) > 0 {
		p = stack[len(stack)-1]
		res = append(res, p.Data)
		//出栈
		stack = stack[:len(stack)-1]
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}
	//此时res保留着前序遍历时,左右互换后的顺序
	front, rear := 0, len(res)-1
	for front < rear {
		tmp := res[rear]
		res[rear] = res[front]
		res[front] = tmp
		front++
		rear--
	}
	return res
}

func LevelOrder(root *BTNode) []string {
	queue := make([]*BTNode, 0)
	res := []string{}
	//求二叉树的宽度,即层数最多的节点
	m := make(map[int]int)
	max := 0
	if root != nil {

		// root入队
		m[1] = 1
		queue = append(queue, root)
		root.Level = 1
		//队列不空
		for len(queue) > 0 {
			//取得队头
			cur := queue[0]
			//出队
			queue = queue[1:]
			res = append(res, cur.Data)
			if cur.Left != nil {
				cur.Left.Level = cur.Level + 1
				m[cur.Left.Level] = m[cur.Left.Level] + 1
				if m[cur.Left.Level] > max {
					max = m[cur.Left.Level]
				}
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Level = cur.Level + 1
				m[cur.Right.Level] = m[cur.Right.Level] + 1
				if m[cur.Right.Level] > max {
					max = m[cur.Right.Level]
				}
				queue = append(queue, cur.Right)
			}
		}
	}
	return res
}

func LevelOrderV2(root *BTNode) [][]string {
	res := make([][]string, 0)
	if root != nil {
		queue := make([]*BTNode, 0)
		//根入队
		queue = append(queue, root)
		//队列不空
		for len(queue) > 0 {
			size := len(queue)
			tmp := make([]string, 0)
			var pre *BTNode
			for i := 0; i < size; i++ {
				// 处理本层节点
				cur := queue[0]
				if pre == nil {
					pre = cur
				} else {
					pre.Next = cur
					pre = cur
				}
				tmp = append(tmp, cur.Data)
				queue = queue[1:]
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
			}
			res = append(res, tmp)
		}
	}
	return res
}

func GetNodeLevel(root *BTNode, value string, level int) int {
	if root == nil {
		return -1
	}
	if root.Data == value {
		return level
	}
	llevel := GetNodeLevel(root.Left, value, level+1)
	if llevel > -1 {
		return llevel
	} else {
		return GetNodeLevel(root.Right, value, level+1)
	}
}

func FindNode(root *BTNode, value string) *BTNode {
	if root == nil {
		return nil
	}
	if root.Data == value {
		return root
	}
	left := FindNode(root.Left, value)
	if left != nil {
		return left
	}
	return FindNode(root.Right, value)
}

func GetBintreeDepth(root *BTNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ldepth := GetBintreeDepth(root.Left)
	rdepth := GetBintreeDepth(root.Right)
	if ldepth > rdepth {
		return ldepth + 1
	} else {
		return rdepth + 1
	}
}

func GetBintreeDepthV2(root *BTNode) int {
	res, depth := 0, 0
	var dfs func(*BTNode)
	dfs = func(root *BTNode) {
		if root != nil {
			depth++
			if root.Left == nil && root.Right == nil {
				//遇到了叶子节点，尝试更新res最大值
				if depth > res {
					res = depth
				}
			}
			dfs(root.Left)
			dfs(root.Right)
			depth--
		}
	}
	dfs(root)
	return res
}

func GetKthNode(root *BTNode, k int) *BTNode {
	var dfs func(*BTNode) *BTNode
	i := 0
	dfs = func(root *BTNode) *BTNode {
		if root != nil {
			i++
			if i == k {
				return root
			}
			left := dfs(root.Left)
			if left != nil {
				return left
			}
			return dfs(root.Right)
		}
		return nil
	}
	return dfs(root)
}
func GetTreeNodeNum(root *BTNode) int {
	num := 0
	var dfs func(*BTNode)
	dfs = func(b *BTNode) {
		if b != nil {
			num++
			dfs(b.Left)
			dfs(b.Right)
		}
	}
	dfs(root)
	return num
}

func GetPreNode(root *BTNode, p *BTNode) *BTNode {
	res := make([]*BTNode, 0)
	var dfs func(*BTNode) *BTNode
	dfs = func(root *BTNode) *BTNode {
		if root != nil {
			res = append(res, root)
			if root == p {
				res = res[:len(res)-1]
				if len(res) > 0 {
					return res[len(res)-1]
				} else {
					return nil
				}
			}
			l := dfs(root.Left)
			if l != nil {
				return l
			} else {
				return dfs(root.Right)
			}
		}
		return nil
	}
	return dfs(root)
}

func GetPreNodev2(root *BTNode, p *BTNode) *BTNode {
	// res := make([]*BTNode, 0)
	var pre *BTNode
	var res *BTNode
	var dfs func(*BTNode)
	dfs = func(c *BTNode) {
		if c != nil {
			if c == p {
				//保存结果
				res = pre
			} else {
				pre = c
			}
			dfs(c.Left)
			dfs(c.Right)

		}
	}
	dfs(root)
	return res
}

//反转一颗二叉树,左右子树互换
func RevertBinTree(root *BTNode) *BTNode {
	var dfs func(*BTNode)
	dfs = func(root *BTNode) {
		if root != nil {
			tmp := root.Left
			root.Left = root.Right
			root.Right = tmp
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return root
}

func RevertBinTreeV2(root *BTNode) *BTNode {
	if root == nil {
		return nil
	}
	left := RevertBinTreeV2(root.Left)
	right := RevertBinTreeV2(root.Right)
	root.Left = right
	root.Right = left
	return root
}

//求某个节点左右子树节点的个数
func GetLeftRightNode(root *BTNode, p *BTNode) []int {
	res := make([]int, 2)
	var dfs func(*BTNode) int
	dfs = func(root *BTNode) int {
		if root != nil {
			l := dfs(root.Left)
			r := dfs(root.Right)
			if root == p {
				res[0] = l
				res[1] = r
			}
			return l + r + 1
		}
		return 0
	}
	dfs(root)
	return res
}

// 求二叉树的最长直径
func GetdiameterOfBinaryTree(root *BTNode) int {
	res := 0
	var dfs func(*BTNode) int
	dfs = func(root *BTNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		cur := l + r
		if cur > res {
			res = cur
		}
		return int(math.Max(float64(l), float64(r))) + 1
	}
	dfs(root)
	return res
}

// 判断B是否为A的子结构,即B的起点为A的根节点
func Recur(A, B *BTNode) bool {
	// 遍历完B
	if B == nil {
		return true
	}
	//A遍历完,B未完,或者A和B的值不相同
	if A == nil || A.Data != B.Data {
		return false
	}
	// A.Data==B.Data
	return Recur(A.Left, B.Left) && Recur(A.Right, B.Right)
}

//判断B是否为A的子结构
func IsSubStructure(A, B *BTNode) bool {
	if A == nil || B == nil {
		return false
	}
	return Recur(A, B) || IsSubStructure(A.Left, B) || IsSubStructure(A.Right, B)
}

func ConstructMaximumBinaryTree(nums []int) *BTNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := 0
	max := nums[0]
	for k, v := range nums {
		if v > max {
			max = v
			maxIndex = k
		}
	}
	root := &BTNode{
		Data: strconv.Itoa(max),
	}
	root.Left = ConstructMaximumBinaryTree(nums[:maxIndex])
	root.Right = ConstructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

func bstFromPreorder(preorder []int) *BTNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &BTNode{
		Data: string(fmt.Sprint(preorder[0])),
	}
	if len(preorder) == 1 {
		return root
	}
	index := 1
	for index < len(preorder) && preorder[index] < preorder[0] {
		index++
	}

	root.Left = bstFromPreorder(preorder[1:index])
	root.Right = bstFromPreorder(preorder[index:])
	return root
}

func GetAllRoute(root *BTNode) [][]string {
	path := make([]string, 0)
	res := make([][]string, 0)
	var dfs func(root *BTNode)
	dfs = func(root *BTNode) {
		if root == nil {
			return
		}
		path = append(path, root.Data)
		if root.Left == nil && root.Right == nil {
			res = append(res, append([]string(nil), path...))

		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}
