/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    flag:=true 
    var Dfs func(*TreeNode, *TreeNode)
    Dfs=func(Left,Right *TreeNode){
        if !flag{
            return 
        }
        if Left==nil && Right==nil{
            flag=true 
            return 
        }
        if Left==nil  || Right==nil {
            flag=false 
            return 
        }
        fmt.Println(Left.Val,Right.Val)
        if Left.Val!=Right.Val{
            flag=false 
            return  
        }

        Dfs(Left.Left,Right.Right)
        Dfs(Left.Right,Right.Left)

    }
    Dfs(root.Left,root.Right)
    if flag{
        return true 
    }
    return  false 
}   

