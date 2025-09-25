func buildArray(nums []int) []int {
    res:=make([]int,len(nums))

    for ind,val := range nums{
        res[ind]=nums[val]
    }
    return res 
}