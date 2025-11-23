func alternatingSum(nums []int) int {
    sum:=0
    for i,val:=range nums{
        if i%2==0{
            sum+=val
        }else{
            sum-=val
        }
    }
    return sum 
}
