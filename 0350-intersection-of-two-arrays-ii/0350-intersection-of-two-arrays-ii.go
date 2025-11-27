func intersect(nums1 []int, nums2 []int) []int {
    firsthave:=make(map[int]int)
    for _,val:=range nums1{
        firsthave[val]++
    }
    res:=[]int{}
    for _,val:=range nums2{
        if firsthave[val]>0{
            firsthave[val]--
            res=append(res,val)
        }
    }
    return res
}