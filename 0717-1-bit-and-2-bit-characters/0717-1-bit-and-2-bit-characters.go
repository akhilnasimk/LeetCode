func isOneBitCharacter(bits []int) bool {
    i:=0
    res:=[][]int{}

    for i<len(bits){
        if bits[i]==0{
            res=append(res,[]int{bits[i]})
        }else if bits[i]==1 {
            res=append(res,[]int{bits[i],bits[i+1]})
            i++
        }
        i++
    }
    fmt.Println(res)
    if len(res[len(res)-1])==1{
        return true 
    }
    return false  
}