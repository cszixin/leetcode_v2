/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-07 10:35:04
 */
package bintree

import (
	"reflect"
	"testing"
)

func TestGetKthNode(t *testing.T) {
	type args struct {
		root *BTNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *BTNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKthNode(tt.args.root, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKthNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTreeNodeNum(t *testing.T) {
	type args struct {
		root *BTNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTreeNodeNum(tt.args.root); got != tt.want {
				t.Errorf("GetTreeNodeNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
