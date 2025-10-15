/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var NewHead *ListNode
    var cur *ListNode 
    carry:=0
    for l1!=nil || l2!=nil || carry>0{
        sum:=carry

        if l1!=nil{
            sum+=l1.Val
            l1=l1.Next
        }
        if l2!=nil{
            sum+=l2.Val
            l2=l2.Next
        }

        N:=&ListNode{Val:sum%10}
        if NewHead ==nil{
            NewHead=N
            cur=N
        }else{
            cur.Next=N
            cur=N
        }
        carry=sum/10
    }
    return NewHead 

}