func nextBeautifulNumber(n int) int {
    res:=0
    i:=n+1
    for {
        maap:=make(map[int]int)
        temp:=i
        for temp>0{
            val:=temp%10
            maap[val]++
            temp=temp/10
        }
        flag:=true 
        for key,val:=range maap{
            if key!=val{
                flag=false 
            }
        }

        if flag{
            res=i
            break
        }else{
            i++
        }
    }
    return res
}