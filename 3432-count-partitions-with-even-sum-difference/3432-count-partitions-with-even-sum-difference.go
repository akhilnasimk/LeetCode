func countPartitions(nums []int) int {
    sum:=0
    for _,val:=range nums{
        sum+=val
    }
    count:=0
    leftsum:=0
    for ind,val:=range nums{
        if ind==len(nums)-1{
            break 
        }
        leftsum+=val
        rightsum:=sum-leftsum
        dif:=leftsum-rightsum
        if dif%2==0{
            fmt.Println()
            count++
        }
    }
    return count  
}