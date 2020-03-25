package leetcode

/**
采用双指针法

设i从头，j从尾
则 S(i, j) = min(height[i], height[j]) * (j - i)

接下来不管是i向后移还是j往前移，水平刻度都会减一，即 (j - i)会变小，
则移动要使 min(height[i], height[j]) 要有增大的可能。
假设：
	长板移动，由于短板不变，则S(i, j) 一定减小
	短板移动，下一个板可能：
		更短： 则 S(i, j) 减小
		一样长： 则 S(i, j) 不变
		更长：则 S(i, j) 变大
所以，短板移动有使得 S(i, j) 变大的可能。
*/
func maxArea(height []int) int {
	maxArea, i, j := 0, 0, len(height)-1
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
