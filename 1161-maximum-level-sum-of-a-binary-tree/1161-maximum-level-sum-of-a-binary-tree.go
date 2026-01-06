/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func maxLevelSum(root *TreeNode) int {
    sumMap:=make(map[int]int)
    var dfs func(*TreeNode,int)
    dfs = func(node *TreeNode ,level int){
        if node==nil{
            return 
        }
        nextSum:=sumMap[level]+node.Val
        sumMap[level]=nextSum
        dfs(node.Left,level+1)
        dfs(node.Right,level+1)
    }    
    dfs(root,1)  
    res:=0
    Biggest:=-10000000
    for ind,val:=range sumMap{
        if val>Biggest{
            res=ind
            Biggest=val
        }
    }
    // fmt.Println(sumMap)
    return res 
}