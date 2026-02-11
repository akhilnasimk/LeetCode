/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    
    sum:=0 
    var dfs func( *TreeNode ,int,bool)

    dfs=func(node *TreeNode ,lastnum int , grand bool ){
        if node==nil{
            return 
        }

        if grand{
            sum+=node.Val
        }
        if lastnum%2==0{
            dfs(node.Left,node.Val,true)
            dfs(node.Right,node.Val,true)
        }else{
            dfs(node.Left,node.Val,false)
            dfs(node.Right,node.Val,false)
        }
    }
    
    dfs(root.Left,root.Val,false)
    dfs(root.Right,root.Val,false)
    return sum 
}