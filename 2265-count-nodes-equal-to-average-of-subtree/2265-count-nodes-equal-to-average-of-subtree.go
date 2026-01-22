/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    count:=0
    var dfs func(*TreeNode )(int,int) 
    dfs=func(node *TreeNode )(int,int) {
        if node==nil{
            return 0,0 
        }
        if node.Left==nil && node.Right==nil{
            count++
            return node.Val,1 
        }
        rsum,rcount:=dfs(node.Right)
        lsum,lcount:=dfs(node.Left)
        sum:=rsum+lsum+node.Val 
        tcount:=rcount+lcount+1
        if sum/tcount==node.Val{
            count++
        }
        return sum,tcount 
    }
    dfs(root)
    return count 
}