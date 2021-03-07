/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l := len(height)
	var max int
	i, j := 0, l-1
	for i < j {
		var area int
		if height[j] > height[i] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}
	return max
}

// @lc code=end

