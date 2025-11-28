func countKDifference(nums []int, k int) int {
    count:=0
    for ind,val:=range nums{
        for i:=ind+1;i<len(nums);i++{
            if Abs(nums[i]-val)==k{
                count++
            }
        }
    }
    return count
}


func Abs(val int)int{
    if val<0{
        return -val
    }
    return val 
}