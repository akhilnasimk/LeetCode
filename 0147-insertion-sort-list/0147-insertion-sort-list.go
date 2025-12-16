/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    dummy := &ListNode{Val: -1}
    cur:=head 
    for cur!=nil{
        next:=cur.Next

        n:=dummy
        for n.Next!=nil&&n.Next.Val<cur.Val{
            n=n.Next 
        }

        nNext:=n.Next 
        n.Next=cur
        cur.Next=nNext 

        cur=next
    }
    return dummy.Next 
}