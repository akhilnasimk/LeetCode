func buildArray(nums []int) []int {
    res:=[] int {}

    for _,val := range nums{
        res=append(res,nums[val])
    }
    return res 
}