func missingNumber(nums []int) int {
    actual:=0
    numLen:=0
    for i:=1 ;i<len(nums)+1;i++{
        actual+=i
        numLen+=nums[i-1]
    }
    return actual-numLen
}