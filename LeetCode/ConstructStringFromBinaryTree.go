package LeetCode

import "strconv"

var	ConstructStringFromBinaryTreeTestTree *TreeNode = &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:2,
		Left:&TreeNode{
			Val:4,
		},
	},
	Right:&TreeNode{
		Val:3,
	},
}

func ConstructStringFromBinaryTree(t *TreeNode) string {
	if t == nil {
		return ""
	}
	s := strconv.Itoa(t.Val)
	if t.Left != nil {
		s = s + "(" + ConstructStringFromBinaryTree( t.Left ) + ")"
	}
	if t.Right != nil {
		if t.Left == nil {
			s = s + "()"
		}
		s = s + "(" + ConstructStringFromBinaryTree( t.Right ) + ")"
	}
	return s
}