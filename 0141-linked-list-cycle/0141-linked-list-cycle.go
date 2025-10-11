/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    cur:=head 
    paths:=[]*ListNode{}
    flag:=false 
    outer:
    for cur!=nil{
        for _,val:=range paths {
            if val==cur {
                flag=true 
                break outer 
            }
        }
        paths=append(paths,cur)
        cur=cur.Next
    }
    if flag{
        return true 
    }
    return false 
}