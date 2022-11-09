/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-26 06:10:50
 * @LastEditors: liuchuan0shi
 * @LastEditTime: 2022-10-19 16:56:11
 */
package main

type MyError struct {
	e error
}

func main() {

	var p *MyError
	var e error = p
	var e1 error = nil
	println(e)
	println(e1)
}
