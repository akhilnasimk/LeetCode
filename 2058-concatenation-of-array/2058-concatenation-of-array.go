func getConcatenation(nums []int) []int {
    length :=len(nums)
    for i:=0 ;i<length;i++{
        nums=append(nums,nums[i])
    }
    return nums
}