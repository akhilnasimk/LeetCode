func addedInteger(nums1 []int, nums2 []int) int {
    firstsum:=0
    for _,val:=range nums1{
        firstsum+=val
    }
    secondsum:=0
    for _,val:=range nums2{
        secondsum+=val
    }

    return (secondsum-firstsum)/len(nums1)
}