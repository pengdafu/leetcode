package leetcode

/**
执行用时 : 24 ms
内存消耗 : 5 MB
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	head := result
	var carry int = 0
	for l1 != nil || l2 != nil {
		l1Data := 0
		l2Data := 0
		if l1 != nil {
			l1Data = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Data = l2.Val
			l2 = l2.Next
		}
		v := (l1Data + l2Data + carry) % 10
		carry = (l1Data + l2Data + carry) / 10
		head.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		head = head.Next
	}
	if carry == 1 {
		head.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return result.Next
}
