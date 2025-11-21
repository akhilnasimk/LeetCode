func findClosest(x int, y int, z int) int {
    xdis:=Abs(x-z)
    ydis:=Abs(y-z)
    if xdis>ydis{
        return 2
    }else if xdis==ydis {
        return 0
    }
    return 1
}

func Abs(val int)int{
    if val<0{
        return -val
    }
    return val 
}