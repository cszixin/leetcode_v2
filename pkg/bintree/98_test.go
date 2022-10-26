/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-19 15:31:18
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 15:46:09
 */
package bintree

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *BTNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
