package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 4, 5, 9}
	target := 8
	if b := TwoSum(nums, target); b[0] != 1 && b[1] != 2 {
		t.Errorf("TwoSum(nums)=%v, want []int{1, 2}", b)
	}
}
