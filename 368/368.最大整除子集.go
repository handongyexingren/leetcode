/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	/*
	sort.Ints(nums)
	n := len(nums)
	visited := make([]int,n)
	for i := 0 ; i < n ; i++ {
		visited[i] = -1
	}
	// 用来记录每一个元素出现过的最靠后的位置
	// 因为我们每次在末尾添加元素时，是判断新元素是否能整除原来的末尾元素
	// 所以，若该元素在前面的情况中，出现在了更靠后的位置，那么我们若在本次将该元素添加到稍靠前的位置一定不能在得到更大的子集了
	// 例如[1,2,3,4,6,8]
	// 前面出现了
	// 1,2,4,8
	// 1,2,6
	// 那么后面诸如1，4，8或者1,3,6就不用再考虑了
	// 因为前面4和6最靠后的位置已经到了下标2，遇到要把4,6放到更靠前位置的情况就可以直接跳过了
	result := make([]int,n)
	length := 0

	var backtrack func(res []int,pos int)
	backtrack = func(res []int,pos int) {
		if len(res) > length {
			length = len(res)
			copy(result,res)
		}
		if len(nums)-pos + len(res) <= length {
			return
		}
		for i := pos ; i < n ; i++ {
			if len(res) == 0 || (len(res) != 0 && nums[i] % res[len(res)-1] == 0) {
				if len(res) <= visited[i] {
					continue
				}
				visited[i] = len(res)
				res = append(res,nums[i])
				backtrack(res,i+1)
				res = res[:len(res)-1]
			}
		}
	}
	backtrack([]int{},0)
	return result[:length]

	*/

	// 看了答案，要用dp啊


	n := len(nums)

	// 记录前一个元素下标
	parent := make([]int,n)
	// 记录子集长度
	dp := make([]int,n)
	dp[0] = 1

	maxindex := 0
	sort.Ints(nums)
	for i := 1 ; i < n ; i++ {
		dp[i] = 1
		parent[i] = i
		for j := i-1 ; j >= 0 ; j-- {
			if nums[i] % nums[j] == 0 {
				if parent[i] == i {
					parent[i] = j
				}else if dp[j] > dp[parent[i]] {
					parent[i] = j
				}
			}
		}
		if parent[i] != -1 {
			dp[i] = dp[parent[i]] + 1
		}

		if dp[i] > dp[maxindex] {
			maxindex = i
		}
	}

	result := make([]int,dp[maxindex])

	for i := len(result)-1 ; i >= 0	; i-- {
		result[i] = nums[maxindex]
		maxindex = parent[maxindex]
	}

	return result
}
// @lc code=end

