func findDifference(nums1 []int, nums2 []int) [][]int {
    m := make(map[int]int) 

    for _, v := range nums1 {
        m[v] |= 1 
    }

    for _, v := range nums2 {
        m[v] |= 2 
    }

    notIn2 := []int{}
    notIn1 := []int{}

    for k, v := range m {
        if v == 1 { 
            notIn2 = append(notIn2, k)
        } else if v == 2 { 
            notIn1 = append(notIn1, k)
        }
    }

    return [][]int{notIn2, notIn1}
}
