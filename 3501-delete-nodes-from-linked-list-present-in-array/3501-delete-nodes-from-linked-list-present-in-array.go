 func modifiedList(nums []int, head *ListNode) *ListNode {
    numMap := make(map[int]bool)
    for _, v := range nums {
        numMap[v] = true
    }

    for head != nil && numMap[head.Val] {
        head = head.Next
    }

    if head == nil {
        return nil
    }

    slow := head
    fast := head.Next

    for fast != nil {
        if numMap[fast.Val] {
            slow.Next = fast.Next
        } else {
            slow = fast
        }
        fast = fast.Next
    }

    return head
}
