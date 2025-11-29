func pairSum(head *ListNode) int {
    // convert to array (very fast)
    arr := []int{}
    for cur := head; cur != nil; cur = cur.Next {
        arr = append(arr, cur.Val)
    }

    // twin pair sum
    n := len(arr)
    maxSum := 0

    for i := 0; i < n/2; i++ {
        s := arr[i] + arr[n-1-i]
        if s > maxSum {
            maxSum = s
        }
    }

    return maxSum
}
