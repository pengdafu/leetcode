package leetcode

/**
题目描述：
给定两个有序数组 nums1 和 nums2。请找出这两个有序数组的中位数。要求时间复杂度为O(log(MIN(m,n)))
example1:
	nums1 = [1, 3], nums2 = [2]   => ans = 2.0
example2:
	num1 = [1, 2], nums2 = [3, 4] => ans = (2 + 3) / 2 = 2.5

分析：
要求时间复杂度为O(log(N))则一定使用二分法，如果没有时间复杂度要求，可以合并一个数组求中位数，时间复杂度为O(m+n)

中位数的定义是在中位数左边的数永远小于等于它，右边的数永远大于等于它。
假设两个数组为A、B，长度分别是m、n。
将数组A分为A[0,i-1] 和 A[i,m]两部分
将数组B分为B[0,j-1] 和 B[j,n]两部分
则 leftPart = { A[0,i-1], B[0,j-1] }, RightPart = { A[i, m], b[j,n] }
则 AB有序数组的中位数一定满足：
	i + j = m - i + n - j (or : m - i + n - j + 1) 如果 n > m, 则 j = (m+n+1)/2 -i(保证j非负)
	A[i-1] <= B[j] && B[j-1] <= A[i]
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1) // 确保 n > m => j > 0
	}
	iMax, iMin, halfLen := m, 0, (m+n+1)/2
	for iMin <= iMax {
		i := (iMax + iMin) / 2
		j := halfLen - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else {
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = Max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}
			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = Min(nums1[i], nums2[j])
			}
			return float64(maxOfLeft+minOfRight) / 2 //注意这里要转float64再除
		}
	}
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
