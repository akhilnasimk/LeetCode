/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    len:=0
    cur:=head
    for cur!=nil{
        len++
        cur=cur.Next 
    }
    i:=1
    first:=head 
    for i<k{
        first=first.Next 
        i++
    } 

    last:=head 
    j:=1 
    for j<=len-k{
        last=last.Next 
        j++
    }

    first.Val,last.Val=last.Val,first.Val 
    // first.Next,last.Next=last.Next,first.Next
    return head 
}