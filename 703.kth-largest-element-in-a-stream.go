package leetcode

type KthLargest struct {
	Heap    []int
	MaxSize int
}

func Constructor(k int, nums []int) KthLargest {
	heap := nums[:k]
	buildK(heap, k)
	for _, v := range nums[k:] {
		if heap[0] < v {
			heap[0] = v
			buildK(heap, k)
		}
	}
	return KthLargest{
		Heap:    heap,
		MaxSize: k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.Heap[0] < val {
		this.Heap[0] = val
		buildK(this.Heap, this.MaxSize)
	}
	return this.Heap[0]
}

func buildK(nums []int, len int) {
	for i := (len - 1) / 2; i >= 0; i++ {
		siftDownK(nums, i, len)
	}
}

func siftDownK(nums []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			return
		}
		if child+1 < hi && nums[child] > nums[child+1] {
			child++
		}
		if nums[root] < nums[child] {
			return
		}
		nums[child], nums[root] = nums[root], nums[child]
		root = child
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
