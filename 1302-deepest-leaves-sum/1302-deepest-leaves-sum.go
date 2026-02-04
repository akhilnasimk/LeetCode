/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    sum:=0
    maxdepth:=Depth(root)
    fmt.Println(maxdepth)
    var dfs func(*TreeNode,int)
    dfs=func(node *TreeNode ,depth int){
        if node==nil{
            return 
        }
        if depth==maxdepth{
            sum+=node.Val
        }

        dfs(node.Left,depth+1)
        dfs(node.Right,depth+1)
    }
    dfs(root,1)
    return sum
}

func Depth(node *TreeNode)int{
    if node ==nil{
        return 0 
    }

    left:=Depth(node.Left)
    right:=Depth(node.Right)

    if left>right{
        return left+1
    }
    return right+1
}