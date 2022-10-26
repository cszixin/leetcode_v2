/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-19 15:31:18
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 16:07:02
 */
package bintree

var pre string = ""

func IsValidBST(root *BTNode) bool {
	if root == nil {
		return true
	} else {
		left := IsValidBST(root.Left)
		if root.Data < pre {
			return false
		}
		pre = root.Data
		right := IsValidBST(root.Right)
		return left && right
	}
}

func isValidBST(root *BTNode) bool {
	var pre *BTNode
	var dfs func(*BTNode) bool
	dfs = func(root *BTNode) bool {

		if root == nil {
			return true
		}

		left := dfs(root.Left)
		if pre != nil && pre.Data >= root.Data {
			return false
		}
		pre = root
		right := dfs(root.Right)
		return left && right
	}
	return dfs(root)
}
