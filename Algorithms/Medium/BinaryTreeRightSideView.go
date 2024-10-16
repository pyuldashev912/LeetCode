package medium

import (
	"container/list"

	"github.com/parviz-yu/LeetCode/Algorithms/helpers"
)

// Time: O(n)
// Space: O(n)
func rightSideView(root *helpers.TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		for range queue.Len() {
			root = queue.Remove(queue.Front()).(*helpers.TreeNode)
			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}
		result = append(result, root.Val)
	}

	return result
}
