package leetcode

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}
	var ans int
	var t = make(map[byte]int)
	for i, j := 0, 0; j < l; j++ {
		if v, ok := t[s[j]]; ok {
			i = max(v+1, i)
		}
		ans = max(ans, j-i+1)
		t[s[j]] = j
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
