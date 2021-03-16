/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				3,
			},
			[]string{"((()))","(()())","(())()","()(())","()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
