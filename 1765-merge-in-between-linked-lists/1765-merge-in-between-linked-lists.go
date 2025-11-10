/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    i:=1
    cur:=list1 
    for i<a {
        cur=cur.Next
        i++
    }
    end:=list1 
    i=0
    for i<=b{
        end=end.Next
        i++
    }
    fmt.Println(end)
    cur.Next=list2
    for cur.Next!=nil{
        cur=cur.Next
    }
    cur.Next=end
    return  list1
}