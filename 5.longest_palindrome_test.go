package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "aba"
	if ans := LongestPalindrome(s); ans != "aba" {
		t.Error("error", ans)
	}
}
