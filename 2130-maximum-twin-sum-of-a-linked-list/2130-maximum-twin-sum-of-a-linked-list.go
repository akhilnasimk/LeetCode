func pairSum(head *ListNode) int {
    length := 0
    for cur := head; cur != nil; cur = cur.Next {
        length++
    }

    index := 0
    twins := make(map[int]int)
    cur := head
    maxSum := 0

    for cur != nil {
        if index < length/2 {
            t := length - index - 1
            twins[t] = cur.Val
        } else {
            twinVal := twins[index]
            s := twinVal + cur.Val
            if s > maxSum {
                maxSum = s
            }
        }
        cur = cur.Next
        index++
    }

    return maxSum
}
