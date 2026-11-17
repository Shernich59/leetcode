/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sum := 0.0
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			sum += float64(curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		avg := sum / float64(size)
		res = append(res, avg)
	}

	return res
}
