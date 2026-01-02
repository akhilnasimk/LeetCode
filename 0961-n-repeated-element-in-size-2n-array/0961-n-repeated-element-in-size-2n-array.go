func repeatedNTimes(nums []int) int {
    n:=len(nums)/2
    mapf:=make(map[int]int)
    res:=0
    for _,val:=range nums{
        mapf[val]++
        if v,_:=mapf[val];v==n{
            res=val
            break 
        }
    }
    return res 
}