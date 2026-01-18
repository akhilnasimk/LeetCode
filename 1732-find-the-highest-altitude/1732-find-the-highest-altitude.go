func largestAltitude(gain []int) int {
    largest:=0
    sum:=0
    for _,val:=range gain{
        sum+=val
        if sum>largest{
            largest=sum
        }
    }
    return largest 
}