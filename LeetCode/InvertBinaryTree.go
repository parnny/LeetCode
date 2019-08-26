package LeetCode

 type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func swapNode(node *TreeNode) {
	if node == nil {
		return
	}
	var tmp *TreeNode
	tmp = node.Left
	node.Left = node.Right
	node.Right = tmp

	swapNode(node.Left)
	swapNode(node.Right)
}

func InvertTree(root *TreeNode) *TreeNode {
	swapNode(root)
	return root
}
