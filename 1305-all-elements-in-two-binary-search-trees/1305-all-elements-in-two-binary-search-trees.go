/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    ar1:=[]int{}
    ar2:=[]int{}
    var dfs func(*TreeNode ,*[]int )
    dfs=func(r1 *TreeNode,arr *[]int ){
        if r1==nil{
            return 
        }
        
        dfs(r1.Left,arr)
        *arr=append(*arr,r1.Val)
        dfs(r1.Right,arr)
    }
    dfs(root1,&ar1)
    dfs(root2,&ar2)
    res := make([]int, 0, len(ar1)+len(ar2))
    i:=0
    j:=0
    for i<len(ar1)&&j<len(ar2){
        if ar1[i]<ar2[j]{
            res=append(res,ar1[i])
            i++
        }else{
            res=append(res,ar2[j])
            j++
        }
    }
    res=append(res,ar1[i:]...)
    res=append(res,ar2[j:]...)
    return res
}