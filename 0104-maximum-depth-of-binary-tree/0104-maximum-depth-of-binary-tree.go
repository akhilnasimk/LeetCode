/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    count:=0

    var dfs func(*TreeNode,int )
    dfs=func(root *TreeNode,c int){
        if root==nil{
            if c>count{
                count=c-1
            }
            return 
        }
        cc:=c+1 
        dfs(root.Left,cc)
        dfs(root.Right,cc)
    }
    dfs(root,1)
    return count
}