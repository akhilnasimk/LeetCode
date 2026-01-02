func repeatedNTimes(nums []int) int {
    mapf:=make(map[int]int)
    for _,val:=range nums{
        mapf[val]++
        if mapf[val]==len(nums)/2{
           return val 
        }
    }
    return -1
}