/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-28 03:13:11
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-28 04:21:03
 */
package array

import (
	"reflect"
	"testing"
)

func TestLongestPalindromeV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test1", args{"babada"}, []string{"bab", "aba", "ada"}},
		{"test2", args{"abcdef"}, []string{"a", "b", "c", "d", "e", "f"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeV2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
