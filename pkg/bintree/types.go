/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:13:18
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-11 11:40:03
 */
package bintree

type BTNode struct {
	Data  string
	Left  *BTNode
	Right *BTNode
	Level int
	Next  *BTNode
}
