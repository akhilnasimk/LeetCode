/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var NHead *ListNode 
    var cur *ListNode 
    curr:=head 
    for curr!=nil{
        if curr.Val !=val {
            node:=&ListNode{Val:curr.Val}
            if NHead==nil{
                cur=node 
                NHead=node
                curr=curr.Next  
            }else{
                cur.Next=node 
                cur=node 
                curr=curr.Next 
            }
        }else{
            curr=curr.Next
        }
    }
    
    return NHead  
}