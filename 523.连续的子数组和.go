/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	/*
	if len(nums) < 2 {
		return false
	}
	sum := make([]int,len(nums)+1)
	summap := make(map[int]int)
	for i := 0 ; i < len(nums) ; i++ {
		sum[i+1] = sum[i] + nums[i]
		summap[sum[i+1]]++
		// 开头有两个0，以及中间有两个0
		if (sum[i+1] == 0 && summap[sum[i+1]] >= 2) || summap[sum[i+1]] >= 3 {
			return true
		}
	}
	for right := 2 ; right < len(sum) ; right++ {
		sumright := sum[right]
		if sumright/k*k == sumright {
			return true
		}
		sumright -= k
		for sumright > 0 {
			if sumright == sum[right-1] {
				// 长度为1，不符合，再减一个
				sumright -= k
				continue
			}
			if summap[sumright] != 0 {
				// 说明有子数组的和为k的倍数。因为我们是通过减了n个k得到的新sumright值
				return true
			}
			sumright -= k
		}
	}
	return false
	*/

	// sum[j]-sum[i] = n*k
	// 那么sum[j]/k - sum[i]/k = n
	// 要使上述减法结果n位整数，那么sum[j],sum[i]分别对k求余的结果相等
	sum := make([]int,len(nums)+1)
	for i := 0 ; i < len(nums) ; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	m := make(map[int]bool)
	for i := 2 ; i < len(sum) ; i++ {
		m[sum[i-2]%k] = true
		if m[sum[i]%k] {
			return true
		}
	}
	return false
}
// @lc code=end

