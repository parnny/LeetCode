package LeetCode

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	result := SearchBST(root.Left,val)
	if result == nil {
		result = SearchBST(root.Right,val)
	}
	return result
}