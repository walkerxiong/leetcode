/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	var (
		minus bool
		start bool
		ret   int
	)
	for _, ch := range []byte(s) {

		if ch == ' ' && !start {
			continue
		}
		if ch == '+' && !start {
			start = true
			continue
		}
		if ch == '-' && !start {
			start = true
			minus = true
			continue
		}
		ch -= '0'
		if ch > 9 {
			break
		}
		ret = ret*10 + int(ch)
		if ret > math.MaxInt32 {
			if minus {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		start = true
	}
	if minus {
		ret = -ret
	}

	return ret
}

// @lc code=end

