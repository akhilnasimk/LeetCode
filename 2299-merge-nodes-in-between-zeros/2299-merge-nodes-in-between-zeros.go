/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    var Nhead *ListNode 
    var Ntale *ListNode 
    cur:=head.Next
    val:=0
    for cur!=nil{
        if cur.Val==0{
            NNode:=&ListNode{Val:val}
            if Nhead==nil {
                Nhead=NNode 
                Ntale=NNode 
            }else{
                Ntale.Next=NNode 
                Ntale=NNode 
            }
            val=0
        }else{
            val+=cur.Val
        }
        cur=cur.Next 
    } 
    return Nhead 
}