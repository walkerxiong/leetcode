/*
 * @lc app=leetcode.cn id=1893 lang=golang
 *
 * [1893] 检查是否区域内所有整数都被覆盖
 */

// @lc code=start
func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		var flag bool
		for _, ran := range ranges {
			if i >= ran[0] && i <= ran[1] {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

// @lc code=end

