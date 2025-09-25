func findDifference(nums1 []int, nums2 []int) [][]int {
    not2:=[]int {}
    not1:=[]int {}
    res:=[][]int {}
    nums1Map:= make(map[int]int);
    for _,val:=range nums1{
        nums1Map[val]++
    }
    nums2Map:= make(map[int]int);
    for _,val:=range nums2{
        nums2Map[val]++
    }
    for key,_ :=range nums1Map{
        if nums2Map[key]==0{
            not2=append(not2,key)
        }
    }
    for key,_ :=range nums2Map{
        if nums1Map[key]==0{
            not1=append(not1,key)
        }
    }
    fmt.Println(not2,not1,nums1Map)
    res=append(res,not2,not1)
    return res

}