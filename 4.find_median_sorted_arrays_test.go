package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	if ans := findMedianSortedArrays(nums1, nums2); ans != 2.5 {
		t.Errorf("findMedianSortedArrays(nums1, nums2) = %v, want %v", ans, 2.5)
	}
}
