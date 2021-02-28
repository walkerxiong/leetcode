/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	l := len(s)
	if l <= numRows || numRows == 1 {
		return s
	}
	var ret []byte
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < l; j += cycleLen {
			ret = append(ret, s[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < l {
				ret = append(ret, s[j+cycleLen-i])
			}
		}
	}
	return string(ret)
}

// @lc code=end

