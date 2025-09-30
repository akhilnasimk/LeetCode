func triangularSum(nums []int) int {
    if len(nums)==1{
        return nums[0]
    }
    newnums := make([]int, len(nums)-1)
    for i := 0; i < len(nums)-1; i++ {
        newnums[i] = (nums[i] + nums[i+1]) % 10
    }
    return triangularSum(newnums)
}