package LeetCode

func IsLeaves(node *TreeNode) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	v := 0
	if root.Left != nil {
		if IsLeaves(root.Left) {
			v = root.Left.Val
		}
	}

	return v + SumOfLeftLeaves(root.Left) + SumOfLeftLeaves(root.Right)
}