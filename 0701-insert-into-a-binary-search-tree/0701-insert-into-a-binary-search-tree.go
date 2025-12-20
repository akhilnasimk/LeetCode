/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root==nil{
        root=&TreeNode{Val:val}
        return root
    }

    var insertion func(*TreeNode)
    temp:=root
    insertion =func(root *TreeNode){
        if root==nil{
            root=&TreeNode{Val:val}
            return 
        }
        if val<root.Val{
            if root.Left==nil{
                root.Left=&TreeNode{Val:val}
                root=nil
                return 
            }
            insertion(root.Left)
        }else{
            if root.Right==nil{
                root.Right=&TreeNode{Val:val}
                root=nil
                return 
            }
            insertion(root.Right)
        }
    }
    insertion(temp)
    return root
}