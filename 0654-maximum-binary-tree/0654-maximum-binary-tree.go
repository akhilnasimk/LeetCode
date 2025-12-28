/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums)==0{
        return nil
    }
    biggest:=0
    index:=0
    for ind,val:=range nums{
        if val>biggest{
            biggest=val
            index=ind
        }
    } 
    Root:=&TreeNode{Val:biggest}
    Root.Left=constructMaximumBinaryTree(nums[:index])
    Root.Right=constructMaximumBinaryTree(nums[index+1:])
    return Root
}

