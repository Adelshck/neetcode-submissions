/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Traversal(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	Traversal(root.Left)
	Traversal(root.Right)
}
func invertTree(root *TreeNode) *TreeNode {
    Traversal(root)
	return root
}
