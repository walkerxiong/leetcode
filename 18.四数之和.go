/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var m = make(map[string]bool)
	l := len(nums)
	if l < 4 {
		return res
	}
	for i := 0; i < l-3; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		for j := i + 1; j < l-2; j++ {
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			left, right := j+1, l-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					key := fmt.Sprintf("%d#%d#%d#%d", nums[i], nums[j], nums[left], nums[right])
					if _, ok := m[key]; !ok {
						m[key] = true
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					}

					left++
				}
				if sum > target {
					right--
				}
				if sum < target {
					left++
				}
			}
		}
	}
	return res
}

// @lc code=end

