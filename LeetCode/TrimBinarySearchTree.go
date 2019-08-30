package LeetCode

var TrimBSTInput1 *TreeNode = &TreeNode{
	Val:3,
	Left:&TreeNode{
		Val:0,
		Right:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:1,
			},
		},
	},
	Right:&TreeNode{
		Val:4,
	},
}


var TrimBSTInput2 *TreeNode = &TreeNode{
	Val:2,
	Left:&TreeNode{
		Val:1,
	},
	Right:&TreeNode{
		Val:3,
	},
}


func TrimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L || root.Val > R {
		var result *TreeNode
		if root.Left != nil {
			result = TrimBST(root.Left,L,R)
		}
		if result != nil {
			return result
		}
		if root.Right != nil {
			result = TrimBST(root.Right,L,R)
		}
		if result != nil {
			return result
		}
		return nil
	} else {
		root.Left = TrimBST(root.Left,L,R)
		root.Right = TrimBST(root.Right,L,R)
		return root
	}
	return nil
}