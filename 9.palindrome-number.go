package leetcode

/**
1、负数肯定不是回文数
2、反转后的数等于原来的数即是回文数
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return _reverse(x) == x
}

func _reverse(x int) int {
	result := 0
	for ; x > 0; x /= 10 {
		result = result*10 + x%10
	}
	return result
}
