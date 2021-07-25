/*
 * @lc app=leetcode.cn id=1743 lang=golang
 *
 * [1743] 从相邻元素对还原数组
 */

// @lc code=start
func restoreArray(adjacentPairs [][]int) []int {
	var (
		m = make(map[int][]int)
		n = len(adjacentPairs) + 1
	)

	for _, nums := range adjacentPairs {
		v := nums[0]
		w := nums[1]
		m[v] = append(m[v], w)
		m[w] = append(m[w], v)
	}

	ans := make([]int, n)
	for k, n := range m {
		if len(n) == 1 {
			ans[0] = k
			break
		}
	}

	ans[1] = m[ans[0]][0]

	for i := 2; i < n; i++ {
		adj := m[ans[i-1]]
		if ans[i-2] == adj[0] {
			ans[i] = adj[1]
		} else {
			ans[i] = adj[0]
		}
	}
	return ans
}

// @lc code=end

