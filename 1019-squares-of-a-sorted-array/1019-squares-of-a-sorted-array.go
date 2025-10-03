func sortedSquares(nums []int) []int {
    for ind,val :=range nums{
        nums[ind]=val*val
    }
    sort.Ints(nums)
    return nums
}