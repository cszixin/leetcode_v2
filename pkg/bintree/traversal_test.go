/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 15:08:30
 */
package bintree

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want *BTNode
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{8, 5, 1, 7, 10, 12}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.args.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllRoute(t *testing.T) {
	type args struct {
		root *BTNode
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllRoute(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
