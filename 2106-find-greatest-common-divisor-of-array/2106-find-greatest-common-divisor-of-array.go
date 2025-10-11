func findGCD(nums []int) int {
    low:=1001
    heigh:=0
    for _,val:=range nums{
        if val<low{
            low=val
        }
        if val>heigh {
            heigh=val
        }
    }
    // fmt.Println(heigh,low)
    divisor:=0
    i:=1
    for i<=heigh{
        if low % i==0 && heigh %i==0{
            divisor=i
        } 
        i++
    }
    // fmt.Println(divisor)
    return divisor
}