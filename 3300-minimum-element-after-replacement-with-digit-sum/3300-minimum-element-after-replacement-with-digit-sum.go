func minElement(nums []int) int {
    min:=100000
    for _,val:=range nums {
        sum:=0
        for val>0{
            sum+=val%10
            val=val/10
        }
        if min>sum{
            min=sum
        }
    }
    return min
}