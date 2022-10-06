/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-06 23:13:41
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
	if root != nil {
		// root入队
		queue = append(queue, root)
		//队列不空
		for len(queue) > 0 {
			//取得队头
			cur := queue[0]
			//处队
			queue = queue[1:]
			res = append(res, cur.Data)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return res
}
