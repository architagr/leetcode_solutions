package average_of_levels_in_binary_tree

import (
	"math"
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

// Cases are the worked examples from the problem statement. The judge
// allows 1e-5, so the comparison does too.
func TestAverageOfLevel(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want []float64
	}{
		{name: "example 1", root: []any{3, 9, 20, nil, nil, 15, 7}, want: []float64{3, 14.5, 11}},
		{name: "example 2", root: []any{3, 9, 20, 15, 7}, want: []float64{3, 14.5, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AverageOfLevel(buildTree(tt.root))
			if len(got) != len(tt.want) {
				t.Fatalf("AverageOfLevel() = %v, want %v", got, tt.want)
			}
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > 1e-5 {
					t.Errorf("AverageOfLevel() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
