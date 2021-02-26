/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	i, j := 0, 0
	var nums = make([]int, 0, l1+l2)
	for i < l1 || j < l2 {
		if i >= l1 && j < l2 {
			nums = append(nums, nums2[j:]...)
			break
		}
		if j >= l2 && i < l1 {
			nums = append(nums, nums1[i:]...)
			break
		}

		if i < l1 && j < l2 {
			if nums1[i] > nums2[j] {
				nums = append(nums, nums2[j])
				j++
			} else {
				nums = append(nums, nums1[i])
				i++
			}
		}
	}
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2
	}
	return float64(nums[int(l/2)])
}

// @lc code=end

