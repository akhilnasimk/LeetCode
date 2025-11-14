func minimumOperations(nums []int) int {
    count:=0
    for _,val:=range nums{
        rem:=val%3
        if rem>0{
            count++
        }
    }
    return count 
}