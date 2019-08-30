package LeetCode

import (
	"strings"
)


type LetterCaseNode struct {
	Number string
	Upper string
	Lower string
	Next *LetterCaseNode
}

func LetterCaseCreateTree(s string) *LetterCaseNode {
	lengh := len(s)
	if lengh == 0 {
		return nil
	}

	node := &LetterCaseNode{}
	i := 0
	c := uint8(0)
	for ;i<lengh; i++ {
		c = s[i]
		if c >= 48 && c <= 57 {
			node.Number+=string(c)
		} else {
			node.Upper = strings.ToUpper(string(c))
			node.Lower = strings.ToLower(string(c))
			break
		}
	}
	if i < lengh {
		node.Next = LetterCaseCreateTree(s[i+1:])
	}
	return node
}

func ConcatLetterCaseNode(pre string, node *LetterCaseNode, list *[]string) {
	if node == nil {
		return
	}
	pre +=node.Number
	if node.Next == nil {
		if node.Lower == node.Upper {
			*list = append(*list, pre)
		} else {
			*list = append(*list, pre+node.Upper)
			*list = append(*list, pre+node.Lower)
		}
	} else {
		ConcatLetterCaseNode( pre+node.Lower, node.Next, list )
		ConcatLetterCaseNode( pre+node.Upper, node.Next, list )
	}
}

func LetterCasePermutation(S string) []string {
	root := LetterCaseCreateTree(S)
	var list []string
	ConcatLetterCaseNode("", root, &list)
	return list
}