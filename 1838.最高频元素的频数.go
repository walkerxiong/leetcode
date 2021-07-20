/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] 最高频元素的频数
 */

// @lc code=start
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	l, r, total := 0, 1, 0
	max := 1

	for ; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		m := r - l + 1
		if m > max {
			max = m
		}
	}
	return max
}

// @lc code=end

