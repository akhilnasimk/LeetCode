/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    var dummy *TreeNode 
    var tail *TreeNode 
    var dfs func( *TreeNode )
    dfs=func(node *TreeNode){
        if node==nil{
            return 
        }

        dfs(node.Left)
        Next:=node.Right 
        node.Left=nil
        node.Right=nil
        if dummy==nil{
            dummy=node 
            tail=node 
        }else{
            tail.Right=node 
            tail=node 
        }
        dfs(Next)
    }

    dfs(root)
    return dummy 
}