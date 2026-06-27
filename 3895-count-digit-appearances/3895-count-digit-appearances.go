func countDigitOccurrences(nums []int, digit int) int {
    total:=0
    for _,val:=range nums{
        total+=ReturnTwoCount(val,digit)
    }
    return total
}

func ReturnTwoCount(num int,digit int )int{
    count:=0
    for num>0{
        rem:=num%10
        // fmt.Println(rem)
        if rem==digit{
            count++
        }
        num=num/10
    }
    return count
}