package LeetCode

import "strconv"

var BinaryTreePathsInput = &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:2,
		Right:&TreeNode{
			Val:5,
		},
	},
	Right:&TreeNode{
		Val:3,
	},
}

func ConcatBinaryTreePath(pre string, node *TreeNode, list *[]string) {
	if node == nil {
		return
	}
	if len(pre) > 0 {
		pre +="->"+strconv.Itoa(node.Val)
	} else {
		pre = strconv.Itoa(node.Val)
	}

	if node.Left == nil && node.Right == nil {
		*list = append(*list, pre)
	} else {
		ConcatBinaryTreePath( pre, node.Left, list )
		ConcatBinaryTreePath( pre, node.Right, list )
	}
}

func BinaryTreePaths(root *TreeNode) []string {
	var result []string
	ConcatBinaryTreePath("",root,&result)
	return result
}