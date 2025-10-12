/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var NHead *ListNode 
    cur :=head 
    for cur!=nil{
        next:=cur.Next 
        cur.Next=NHead
        NHead=cur 
        cur=next 
    }
    return NHead
}