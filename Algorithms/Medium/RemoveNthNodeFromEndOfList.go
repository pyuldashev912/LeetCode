package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func removeNthFromEnd(head *helpers.ListNode, n int) *helpers.ListNode {
	fast, slow := head, head
	for ; n > 0; n-- {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}
