package findacorrespondingnodeofabinarytreeinclone

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func getTargetCopy(original, clones, target *TreeNode) *TreeNode {
	if clones == nil {
		return nil
	}
	if clones.Val == target.Val {
		return clones
	}
	n := getTargetCopy(original.Left, clones.Left, target)
	if n != nil {
		return n
	}
	return getTargetCopy(original.Right, clones.Right, target)
}
