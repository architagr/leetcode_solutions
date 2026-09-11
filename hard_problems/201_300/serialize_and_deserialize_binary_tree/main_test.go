package serializeanddeserializebinarytree

import (
	"reflect"
	"testing"
)

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

func treeVals(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}
	out := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

// The encoding is the codec's own business, so the case checks the round
// trip rather than the string: deserialize(serialize(t)) must give back
// the tree that went in.
func TestCodecRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		root []any
	}{
		{name: "example 1", root: []any{1, 2, 3, nil, nil, 4, 5}},
		{name: "example 2 - empty", root: []any{}},
		{name: "single node", root: []any{7}},
		{name: "left leaning", root: []any{1, 2, nil, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ser := Constructor()
			deser := Constructor()
			got := treeVals(deser.deserialize(ser.serialize(buildTree(tt.root))))
			want := treeVals(buildTree(tt.root))
			if !reflect.DeepEqual(got, want) {
				t.Errorf("round trip = %v, want %v", got, want)
			}
		})
	}
}
