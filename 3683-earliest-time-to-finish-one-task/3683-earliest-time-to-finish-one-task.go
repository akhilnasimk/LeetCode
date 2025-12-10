func earliestTime(tasks [][]int) int {
    fastest:=2001
    for _,val:=range tasks{
        sum:=val[0]+val[1]
        if sum<fastest{
            fastest=sum
        }
    }
    return fastest
}