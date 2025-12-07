func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    map1 := make(map[int]int)
    keys := []int{}

    for _, val := range nums1 {
        map1[val[0]] = val[1]
        keys = append(keys, val[0])
    }

    for _, val := range nums2 {
        v, exist := map1[val[0]]
        if exist {
            map1[val[0]] = v + val[1]
        } else {
            map1[val[0]] = val[1]
            keys = append(keys, val[0])
        }
    }

    // Remove duplicate keys before sorting
    keys = unique(keys)
    sort.Ints(keys)

    res := [][]int{}
    for _, key := range keys {
        res = append(res, []int{key, map1[key]})
    }

    return res
}

func unique(nums []int) []int {
    seen := map[int]bool{}
    res := []int{}
    for _, n := range nums {
        if !seen[n] {
            seen[n] = true
            res = append(res, n)
        }
    }
    return res
}
