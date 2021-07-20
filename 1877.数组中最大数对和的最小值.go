/*
 * @lc app=leetcode.cn id=1877 lang=golang
 *
 * [1877] 数组中最大数对和的最小值
 */

// @lc code=start
func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for k, v := range nums[:n/2] {
		sum := nums[n-k-1] + v
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

// @lc code=end

