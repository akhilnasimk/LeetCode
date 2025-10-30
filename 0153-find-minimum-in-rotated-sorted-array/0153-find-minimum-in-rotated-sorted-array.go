func findMin(nums []int) int {
    min:=10000
    for _,val:=range nums{
        if min>val{
            min=val
        }
    }
    return min
}