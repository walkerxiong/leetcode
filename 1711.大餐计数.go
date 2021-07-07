/*
 * @lc app=leetcode.cn id=1711 lang=golang
 *
 * [1711] 大餐计数
 */

// @lc code=start
func countPairs(deliciousness []int) int {
	var maxVal = deliciousness[0]
	for _, v := range deliciousness[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	maxSum := maxVal * 2
	var ans int
	var cnt = map[int]int{}
	for _, v := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-v]
		}
		cnt[v]++
	}
	return ans % (1e9 + 7)
}

// @lc code=end

