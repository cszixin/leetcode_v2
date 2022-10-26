/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-17 10:24:29
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 10:32:32
 */
package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test1", args{9}, [][]int{{2, 3, 4}, {4, 5}}},
		{"test2", args{15}, [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
