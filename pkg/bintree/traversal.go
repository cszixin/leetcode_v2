/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-07 20:10:30
 */
package bintree

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

func Postrder(root *BTNode) []string {
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
