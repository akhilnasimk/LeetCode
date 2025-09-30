func maximum69Number (num int) int {
    sl:= []int{}
    temp:=num
    for temp>0{
        sl=append(sl,temp%10)
        temp=temp/10
    }
    for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
    sl[i], sl[j] = sl[j], sl[i]
}
    great:=0
    change:=0
    for change<len(sl){
        num:=0
        for i:=0;i<len(sl);i++{
            if i==change{
                if sl[i]==6{
                    num=num*10 +9
                }else{
                    num=num*10 +6
                }
            }else{
                num=num*10+sl[i]
            }
        }
        fmt.Println(num)
        change++
        if num>great{
            great=num
        }
    }
    fmt.Println(sl)
    if great>num{
        return great 
    }
    return num
}