func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    odd := head
    cur := head.Next
    i := 1 

    for cur != nil && cur.Next != nil {
        if i%2 == 1 { 
            old := cur.Next        
            oddold := odd.Next    
            cur.Next = old.Next
            odd.Next = old
            old.Next = oddold
            odd = old
        } else {
            cur = cur.Next
        }

        i++
    }

    return head
}
