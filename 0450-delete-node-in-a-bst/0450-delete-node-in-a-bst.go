/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root==nil{
        return nil
    }

    if root.Val<key{
        root.Right=deleteNode(root.Right,key)
    }else if root.Val>key{
        root.Left=deleteNode(root.Left,key)
    }else{
        if root.Right==nil && root.Left==nil{
            root=nil
        }else if root.Right==nil{
            root=root.Left
        }else if root.Left==nil{
            root=root.Right
        }else{
            min:=MinFinder(root.Right)
            root.Val=min
            root.Right=deleteNode(root.Right,min)
        }
    }

    return root 
}

//finiding smallest on the node after node to delete
func MinFinder(R *TreeNode) int {
	if R.Left == nil {
		return R.Val
	}
	return MinFinder(R.Left)
}