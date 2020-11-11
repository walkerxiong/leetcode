package main

func lengthOfLongestSubstring(s string) int {
	var (
		l      = len(s)
		maxLen int
	)
	if l <= 1 {
		return l
	}
	for i := 0; i < l; i++ {
		var m = make(map[byte]bool)
		for j := i; j < l; j++ {
			bt := s[j]
			if _, ok := m[bt]; ok {
				break
			}
			m[bt] = true
		}
		mL := len(m)
		if mL >= maxLen {
			maxLen = mL
		}

	}
	return maxLen
}
