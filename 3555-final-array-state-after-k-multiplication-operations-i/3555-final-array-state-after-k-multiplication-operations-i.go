func getFinalState(nums []int, k int, multiplier int) []int {
    i:=0
    for i<k{
        low:=100000
        lowind:=0
        for i,v:=range nums{
            if low>v{
                low=v
                lowind=i
            }
        }
        nums[lowind]=nums[lowind]*multiplier;
        i++
    }
    return nums
}