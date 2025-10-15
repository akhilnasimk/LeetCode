/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    len:=0
    cur:=head 
    for cur!=nil{
        cur=cur.Next
        len++
    }
    if len==1{
        head=nil
        return head
    }
    fmt.Println(len)
    togo:=len-n
    if togo==0{
        head=head.Next
        return head 
    }
    fmt.Println(togo)
    start:=1
    a:=head
    for start<togo{
        a=a.Next 
        start++
    }
    if a.Next.Next==nil{
        a.Next=nil
        return head 
    }
    a.Next=a.Next.Next 
    return head 
}