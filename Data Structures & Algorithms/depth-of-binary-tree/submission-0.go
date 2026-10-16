/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	
    if root == nil {
		return 0
	}
	count := 1
	count += max(maxDepth(root.Right), maxDepth(root.Left))
	return count
}
