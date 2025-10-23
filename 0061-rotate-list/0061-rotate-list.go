/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k==0 || head==nil || head.Next==nil{
        return head 
    }
    C:=head 
    len:=0
    for C!=nil{
        C=C.Next
        len++
    }
    i:=1
    for i<=k%len{
        cur:=head
        for cur.Next.Next!=nil{
            cur=cur.Next
        }
        cur.Next.Next=head 
        head=cur.Next
        cur.Next=nil
        // fmt.Println(cur)
        i++
    }
    return head 
}