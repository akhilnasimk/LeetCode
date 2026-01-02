func repeatedNTimes(nums []int) int {
    n:=len(nums)/2
    mapf:=make(map[int]int)
    for _,val:=range nums{
        mapf[val]++
        if v,_:=mapf[val];v==n{
           return val 
        }
    }
    return -1
}