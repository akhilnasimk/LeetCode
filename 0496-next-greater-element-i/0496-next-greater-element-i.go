func nextGreaterElement(nums1 []int, nums2 []int) []int {
    res:=[]int{}
    nums2map:= make(map[int]int)
    for ind,val := range nums2{
        nums2map[val]=ind
    }
    fmt.Println(nums2map)
    for _,val:=range nums1{
        ind:=nums2map[val]
        flag:=true 
        for ind<len(nums2){
            if nums2[ind]>val{
                res= append(res,nums2[ind])
                flag=false 
                break 
            }
            ind++
        }
        if flag{
            res=append(res,-1)
        }
        
    }
    return res
}