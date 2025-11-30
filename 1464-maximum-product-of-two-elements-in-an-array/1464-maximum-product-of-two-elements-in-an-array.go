func maxProduct(nums []int) int {
    biggest:=0
    second:=0
    for _,val:=range nums{
        if val>=biggest{
            second=biggest 
            biggest=val
        }else if val>second{
            second=val
        }
    }
    return (biggest-1) * (second-1)
}