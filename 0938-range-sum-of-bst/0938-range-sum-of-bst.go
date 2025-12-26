/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum:=0
    var dfs func(*TreeNode)
    dfs=func(root *TreeNode ){
        if root==nil{
            return 
        }
        if root.Val>=low && root.Val<=high{
            sum+=root.Val
        }

        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return sum 
}