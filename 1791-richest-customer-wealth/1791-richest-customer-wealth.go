func maximumWealth(accounts [][]int) int {
    rich:=0
    for _,val :=range accounts{
        sum:=0
        for _,num :=range val{
            sum+=num
        }
        if sum>rich{
            rich=sum
        }
    }
    return rich
}