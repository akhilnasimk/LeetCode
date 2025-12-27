/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type MinHeap []int

func (R MinHeap)Len()int{
    return len(R)
}

func (R MinHeap)Swap(i,j int ){
    R[i],R[j]=R[j],R[i]
}

func (R MinHeap)Less(i,j int )bool{
    return R[i]<R[j]
}

func (R *MinHeap)Push(val interface{}){
    *R=append(*R,val.(int))
}

func (R *MinHeap)Pop()interface{}{
    old:=*R
    n:=old.Len()
    res:=old[n-1]
    *R=old[:n-1]
    return res 
}

func kthSmallest(root *TreeNode, k int) int {
    h:=&MinHeap{}
    heap.Init(h)

    Dfs(root,h)
    var ans int
    for i := 0; i < k; i++ {
        ans=heap.Pop(h).(int)
    }

    return ans
}

func Dfs(root *TreeNode,minH *MinHeap){
    if root==nil{
        return 
    }
    heap.Push(minH,root.Val)
    Dfs(root.Left,minH)
    Dfs(root.Right,minH)
}