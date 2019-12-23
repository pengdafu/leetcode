package leetcode

func TwoSum(nums []int, target int) []int {
	// nums 为 nil 或者长度小于2 都不可能存在target
	if nums == nil || len(nums) < 2 {
		return nil
	}
	// value: index
	pc := make(map[int]int)
	for i, v := range nums {
		if ri, ok := pc[target-v]; ok {
			return []int{ri, i}
		}
		pc[v] = i
	}
	return nil
}
