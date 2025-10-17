/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    var Head *ListNode 
    var tail *ListNode 
    if list1 ==nil {
        Head= list2
        return Head
    }
    if list2==nil{
        Head=list1
        return Head
    }
    cur1:=list1
    cur2:=list2

    for cur1!=nil && cur2!=nil{
        if cur1.Val <= cur2.Val{
            if Head==nil{
                Head=cur1 
                tail=cur1
                cur1=cur1.Next
            }else{
                tail.Next=cur1
                tail=cur1
                cur1=cur1.Next
            }
        }else if cur2.Val < cur1.Val {
            if Head==nil {
                Head=cur2 
                tail=cur2
                cur2=cur2.Next
            }else{
                tail.Next=cur2
                tail=cur2
                cur2=cur2.Next
            }
        }
    }
    fmt.Println(tail,Head)
    fmt.Println(cur1,cur2)
    if cur1!=nil{
        tail.Next=cur1
    }
    if cur2!=nil{
        tail.Next=cur2
    }
    return Head  
}