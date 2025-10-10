func differenceOfSum(nums []int) int {
    digisum:=0
    sum:=0
    for _,val :=range nums{
        sum+=val
        isum:=0
        for val>0{
            isum+=val%10
            val=val/10
        }
        digisum+=isum
    }
    return int(math.Abs(float64(digisum)-float64(sum)))
}