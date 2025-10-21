/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head==nil || head.Next==nil{
        return head
    }
    cur:=head
    fast:=cur.Next 
    for fast!=nil{
        if cur==head {
            head=cur.Next 
        }
        old:=fast.Next 
        fast.Next=cur
        if old!=nil{
            if old.Next!=nil{
                cur.Next=old.Next
            }else{
                cur.Next=old
            }
        }else{
            cur.Next=nil
        }
        
        if old!=nil{
            cur=old
            fast=old.Next
        }else{
            fast=nil
        }
    }
    return head 
}