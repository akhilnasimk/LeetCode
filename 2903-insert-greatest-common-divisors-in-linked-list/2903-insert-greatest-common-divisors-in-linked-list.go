/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    cur:=head 
    for cur.Next!=nil{
        N1:=cur.Val
        N2:=cur.Next.Val
        divisor:=0
        i:=1
        small:=Smaller(N1,N2)
        for i <= small{
            if N1%i==0 && N2%i==0{
                if i>divisor{
                    divisor=i
                }
            }
            i++
        }
        node:=&ListNode{Val:divisor}
        temp:=cur.Next
        cur.Next=node 
        node.Next=temp 
        cur=cur.Next.Next 
    }
    return head 
}


func Smaller(a,b int )int{
    if a <b {
        return a 
    }
    return b 
}