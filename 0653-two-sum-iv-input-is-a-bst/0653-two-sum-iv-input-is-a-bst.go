/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    element:=make(map[int]bool )
    flag:=false 
    var dfs func( *TreeNode)
    dfs=func(node *TreeNode){
        if node==nil{
            return 
        }
        if element[k-node.Val]{
            flag=true 
            return 
        }
        element[node.Val]=true 
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    if flag{
        return true 
    }
    return false 
}