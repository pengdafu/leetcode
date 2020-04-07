package leetcode

/**
解法1：全部排序后取第K大
*/
import (
	"sort"
)

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/*
解法二：维护K大小的堆
*/
func FindKthLargest2(nums []int, k int) int {
	heap := nums[:k]
	build(heap, k)
	for _, v := range nums[k:] {
		if heap[0] < v {
			heap[0] = v
			build(heap, k)
		}
	}
	return heap[0]
}

func build(nums []int, len int) {
	for i := (len - 1) / 2; i >= 0; i-- {
		siftDown(nums, i, len, 0)
	}
}

func siftDown(nums []int, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && nums[child] > nums[child+1] {
			child++
		}
		if nums[root] < nums[child] {
			return
		}
		swap(nums, root, child)
		root = child
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
