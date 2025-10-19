/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode ) *ListNode {
    if head==nil || head.Next==nil{
        return head 
    }
    cur:=head 
    last:=head.Next

    for last!=nil{
        if cur.Val==last.Val{
            last=last.Next 
        }else if cur.Val<last.Val{
            cur.Next=last
            cur=last
            last=last.Next
        }
    } 
    cur.Next=nil
    return head 
}