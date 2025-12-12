func countDistinctIntegers(nums []int) int {
    unique:=make(map[int]bool)
    count:=0
    for i:=range len(nums){
        _,yee:=unique[nums[i]]
        if !yee{
            // fmt.Println("from array",nums[i])
            unique[nums[i]]=true
            count++
        }
        temp:=nums[i]/10
        newnum:=nums[i]%10
        for temp>0{
            rem:=temp%10
            newnum=(10*newnum)+rem
            temp=temp/10
        }
        _,exist:=unique[newnum]
        if !exist{
            unique[newnum]=true
            // fmt.Println(newnum)
            count++
        }
    }
    return count
}