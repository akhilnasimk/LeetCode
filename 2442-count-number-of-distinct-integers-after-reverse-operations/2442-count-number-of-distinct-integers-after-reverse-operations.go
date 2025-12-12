func countDistinctIntegers(nums []int) int {
    unique:=make(map[int]bool)
    for i:=range len(nums){
        unique[nums[i]]=true
        temp:=nums[i]/10
        newnum:=nums[i]%10
        for temp>0{
            rem:=temp%10
            newnum=(10*newnum)+rem
            temp=temp/10
        }
        unique[newnum]=true
    }
    return len(unique)
}