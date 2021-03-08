/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	var (
		sum    = 0
		inited bool
		diff   float64
	)
	for start := 0; start < l; start++ {
		i, j := start+1, l-1
		for i < j {
			s := nums[start] + nums[i] + nums[j]

			if s == target {
				return s
			}
			d := math.Abs(float64(s - target))
			if !inited {
				diff = d
				inited = true
				sum = s
			}
			if d < diff {
				diff = d
				sum = s
			}
			if s > target {
				if nums[i] > nums[j] {
					i++
				} else {
					j--
				}
			} else {
				if nums[i] > nums[j] {
					j--
				} else {
					i++
				}
			}
		}
	}
	return sum
}

// @lc code=end

