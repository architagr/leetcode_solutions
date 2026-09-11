package construct_string_from_binary_tree

import "testing"

// buildTree rebuilds a tree from the level-order form the examples use,
// where nil marks a missing child.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestTree2str(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want string
	}{
		{name: "example 1", root: []any{1, 2, 3, 4}, want: "1(2(4))(3)"},
		{name: "example 2", root: []any{1, 2, 3, nil, 4}, want: "1(2()(4))(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tree2str(buildTree(tt.root)); got != tt.want {
				t.Errorf("Tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
