/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-19 12:57:25
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 13:01:37
 */
package bintree

func flatten(root *BTNode) {
	var dfs func(root *BTNode) *BTNode
	dfs = func(root *BTNode) *BTNode {
		if root == nil {
			return nil
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		p := left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = right
		root.Right = left
		return left
	}
	dfs(root)
}
