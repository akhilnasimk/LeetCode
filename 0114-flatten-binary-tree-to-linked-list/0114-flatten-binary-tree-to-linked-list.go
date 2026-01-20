/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)*TreeNode   {
    var head *TreeNode 
    var tail *TreeNode

    var dfs func(*TreeNode)
    dfs=func(node *TreeNode ){
        if node==nil{
            return 
        }
        left:=node.Left 
        right:=node.Right
        node.Left=nil
        node.Left=nil
        if head==nil && tail==nil{
            head=node 
            tail=node 
        }else{
            tail.Right=node
            tail=node
        }
        dfs(left)
        dfs(right)
    }
    dfs(root)
    return head 
}