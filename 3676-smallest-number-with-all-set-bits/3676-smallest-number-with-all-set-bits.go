func smallestNumber(n int) int {
    i:=n
    for{
        flag:=true 
        temp:=i
        for temp>0{
            rem:=temp%2
            if rem==0{
                flag=false 
                break 
            }
            temp=temp/2
        }
        if flag {
            break 
        }
        i++

    }
    return i
}