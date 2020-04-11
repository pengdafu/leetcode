package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}
	res, windows := []int{}, []int{}
	for i, num := range nums {
		if i >= k && i-windows[0] >= k {
			windows = windows[1:]
		}
		// 注意这里的windows是一直在变的，length也一直在变
		for len(windows) > 0 && nums[windows[len(windows)-1]] <= num {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i)
		// 当已经满足一个窗口长度的时候，才增加结果值
		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}
