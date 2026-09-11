package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ss []string
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ss = append(ss, "")
			return
		}
		ss = append(ss, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return strings.Join(ss, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")

	var dfs func(ss []string) ([]string, *TreeNode)
	dfs = func(ss []string) ([]string, *TreeNode) {
		if ss[0] == "" {
			return ss[1:], nil
		}

		v, _ := strconv.Atoi(ss[0])
		root := &TreeNode{
			Val: v,
		}

		ss, root.Left = dfs(ss[1:])
		ss, root.Right = dfs(ss)

		return ss, root
	}

	_, root := dfs(ss)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
