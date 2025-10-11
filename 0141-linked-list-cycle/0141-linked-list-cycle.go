func hasCycle(head *ListNode) bool {
    cur := head
    paths := []*ListNode{}

    for cur != nil {
        for _, node := range paths {
            if node == cur {
                return true // directly return
            }
        }
        paths = append(paths, cur)
        cur = cur.Next
    }

    return false
}