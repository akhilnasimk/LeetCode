func minMoves(nums []int) int {
    biggest:=0
    for _,val:=range nums{
        if val>biggest{
            biggest=val
        }
    }
    res:=0
    for _,val:=range nums{
        res+=biggest- val 
    }
    return res 
}