package LeetCode

var CousinsInBinaryTreeInput1 = &TreeNode{
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

var CousinsInBinaryTreeInput2 = &TreeNode{
	Val:1,
	Left:&TreeNode{
		Val:2,
		Right:&TreeNode{
			Val:4,
		},
	},
	Right:&TreeNode{
		Val:3,
		Right:&TreeNode{
			Val:5,
		},
	},
}

func FindInBinaryTree(root *TreeNode, value int, deep int) (int, *TreeNode) {	// return (deep, )
	if root == nil {
		return 0,nil
	}
	if root.Left != nil {
		if root.Left.Val == value {
			return deep,root
		}
		rd,rp := FindInBinaryTree(root.Left,value,deep+1)
		if rp != nil {
			return rd,rp
		}
	}
	if root.Right != nil {
		if root.Right.Val == value {
			return deep, root
		}
		rd,rp :=  FindInBinaryTree(root.Right,value,deep+1)
		if rp != nil {
			return rd,rp
		}
	}
	return 0,nil
}

func IsCousinsInBinaryTree(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	deep1,parent1 := FindInBinaryTree(root,x,0)
	deep2,parent2 := FindInBinaryTree(root,y,0)
	return (deep1 == deep2) && (parent1 != parent2) && (parent1 != nil) && (parent2 != nil)
}