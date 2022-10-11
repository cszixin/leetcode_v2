/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-06 21:31:00
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-10 16:12:33
 */
package bintree

import (
	"reflect"
	"testing"
)

func TestGetdiameterOfBinaryTree(t *testing.T) {
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
			if got := GetdiameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("GetdiameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevelOrderV2(t *testing.T) {
	type args struct {
		root *BTNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrderV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertBinTreeV2(t *testing.T) {
	type args struct {
		root *BTNode
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
			if got := RevertBinTreeV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevertBinTreeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
