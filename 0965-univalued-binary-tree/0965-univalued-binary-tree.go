/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root==nil{
        return true 
    }
    var dfs func (*TreeNode ,int )bool
    dfs=func(root *TreeNode ,last int)bool{
        if root==nil{
            return true 
        }
        if root.Val!=last{
            return false 
        }
        
        return dfs(root.Right,root.Val)&& dfs(root.Left,root.Val)
    }

    return dfs(root,root.Val)
}

