func sumFourDivisors(nums []int) int {
    sum:=0
    for _,val:=range nums{
        sum+=Sum(val)
    }

    return sum 
}


func Sum(val int )int {
    sum:=val
    count:=1
    for i:=1 ;i<=val/2;i++{
        if val%i==0{
            count++
            sum+=i
        }

        if count>4{
            return 0 
        }
    }

    if count ==4{
        return sum
    }
    return 0
}