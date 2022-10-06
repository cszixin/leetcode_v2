/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:15:24
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-06 21:36:14
 */
package bintree

func CreateBinTree(preorder, inorder []string) *BTNode {
	if len(preorder) < 1 {
		return nil
	}
	root := &BTNode{
		Data: preorder[0],
	}
	//找到在inorder中的index
	index := -1
	for k, v := range inorder {
		if root.Data == v {
			index = k
		}
	}
	preLeft, preRight := preorder[1:index+1], preorder[index+1:]
	inLeft, inRight := inorder[:index], inorder[index+1:]
	root.Left = CreateBinTree(preLeft, inLeft)
	root.Right = CreateBinTree(preRight, inRight)
	return root
}
