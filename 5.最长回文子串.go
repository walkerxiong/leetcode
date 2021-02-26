/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	ans := ""
	longest := 0
	for i := 0; i < l; i++ {
		for j := i + longest; j <= l; j++ {
			sub := s[i:j]
			if isPalindrom(sub) && longest < len(sub) {
				longest = len(sub)
				ans = sub
			}
		}
	}
	return ans
}

func isPalindrom(s string) bool {
	l := len(s)
	if l <= 1 {
		return true
	}
	mid := 0
	if l%2 == 0 {
		mid = l / 2
	} else {
		mid = l/2 + 1
	}

	for i := mid; i < l; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true

}

// @lc code=end

