func findFinalValue(nums []int, original int) int {
    collection:=make(map[int]bool)
    for _,val:=range nums{
        collection[val]=true
    }
    for{
        _,exist:=collection[original]
        if !exist{
            break 
        }
        original*=2 
    }
    return original 
}