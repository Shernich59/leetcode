/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DFS(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return ((left.Val == right.Val) && DFS(left.Left, right.Right) && DFS(left.Right, right.Left))
}

func isSymmetric(root *TreeNode) bool {
	return DFS(root.Left, root.Right)
}