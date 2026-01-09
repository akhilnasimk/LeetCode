/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    num:=0
    len:=0
    var godeep func (*ListNode )
    godeep=func(node *ListNode ){
        if node==nil{
            return 
        }
        godeep(node.Next)
        num+=node.Val*(1<<len)
        len++
    }
    godeep(head)
    return num 
}