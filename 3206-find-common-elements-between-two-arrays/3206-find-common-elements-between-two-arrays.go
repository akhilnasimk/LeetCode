func findIntersectionValues(nums1 []int, nums2 []int) []int {
    ans1:=0
    ans2:=0
    mapp1:=make(map[int]int)
    for _,val:=range nums1{
        mapp1[val]++
    }
    mapp2:=make(map[int]int)
    for _,val:=range nums2{
        mapp2[val]++
    }
    for _,val:=range nums1{
        if mapp2[val]!=0{
            ans1++
        }
    }
    for _,val:=range nums2{
        if mapp1[val]!=0{
            ans2++
        }
    }
    fmt.Println(mapp1,mapp2)
    return []int{ans1,ans2}
}