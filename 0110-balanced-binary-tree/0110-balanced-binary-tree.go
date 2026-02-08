/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return Depth(root) != -1
}

func Depth(node *TreeNode) int {
    if node == nil {
        return 0
    }

    left := Depth(node.Left)
    if left == -1 {
        return -1
    }

    right := Depth(node.Right)
    if right == -1 {
        return -1
    }

    if Abs(left-right) > 1 {
        return -1
    }

    if left > right {
        return left + 1
    }

    return right + 1
}

func Abs(val int) int {
    if val < 0 {
        return -val
    }
    return val
}
