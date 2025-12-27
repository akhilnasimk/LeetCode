/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    dep:=Depth(root)
    res := make([][]int, dep)

    Add(root,0,&res)

    return res

}

func Add(root *TreeNode ,index int , array *[][]int){
    if root==nil{
        return 
    }
    (*array)[index]=append((*array)[index],root.Val)
    Add(root.Left,index+1,array)
    Add(root.Right,index+1,array)
}

func Depth(root *TreeNode)int{
    if root==nil{
        return 0 
    }

    left:=Depth(root.Left)
    right:=Depth(root.Right)

    if left>right{
        return left+1
    }
    return right+1
}