/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-28 03:13:11
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-28 04:23:55
 */
package test

import (
	"leetcode/pkg/array"
	"reflect"
	"testing"
)

func TestIsPalinDrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{s: "abc"}, false},
		{"test2", args{s: "a"}, true},
		{"test3", args{s: ""}, true},
		{"test4", args{s: "abba"}, true},
		{"test5", args{s: "abcba"}, true},
		{"test6", args{s: "abccba"}, true},
		{"test7", args{s: "aabccba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.IsPalinDrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalinDrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPalinDrome(t *testing.T) {
	type args struct {
		s     string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"abbac", 1, 2}, "abba"},
		{"test2", args{"ababac", 2, 2}, "ababa"},
		{"test3", args{"ababac", 0, 0}, "a"},
		{"test4", args{"ababac", 0, 1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.GetPalinDrome(tt.args.s, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("GetPalinDrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"babad"}, "bab"},
		{"test2", args{"cbbd"}, "bb"},
		{"test3", args{"abba"}, "abba"},
		{"test4", args{"abbcba"}, "bcb"},
		{"test5", args{"abhecda"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"test3", args{"cbbd"}, []string{"bb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := array.LongestPalindromeV2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
