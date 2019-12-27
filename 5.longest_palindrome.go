package leetcode

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

输入: "cbbd"
输出: "bb"

*/
// 动态规划解法
func LongestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	length := len(s)
	var left, right int
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	for i := 1; i < length; i++ {
		dp[i][i] = true
		for j := i - 1; j >= 0; j-- {
			dp[j][i] = s[i] == s[j] && (i-j < 3 || dp[j+1][i-1])
			if dp[j][i] && (i-j) > (right-left) {
				right, left = i, j
			}
		}
	}
	return s[left : right+1]
}

// 中心扩展法

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var maxExpand string
	length := len(s)
	var expandFunc = func(mid int) string {
		i, j := mid-1, mid+1
		for j < length && s[j] == s[mid] {
			j++
		}
		for i >= 0 && j < length && s[i] == s[j] {
			i--
			j++
		}
		return s[i+1 : j]
	}
	for i := 0; i < length-1; i++ {
		l := expandFunc(i)
		if len(l) > len(maxExpand) {
			maxExpand = l
		}
	}
	return maxExpand
}
