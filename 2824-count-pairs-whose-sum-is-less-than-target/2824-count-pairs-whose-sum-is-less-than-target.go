func countPairs(nums []int, target int) int {
    count:=0
    for ind,val:=range nums{
        for j:=ind+1;j<len(nums);j++{
            if ind<j && val+nums[j]<target{
                count++
            }
        }
    }
    return count 
}