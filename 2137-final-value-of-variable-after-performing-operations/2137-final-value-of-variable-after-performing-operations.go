func finalValueAfterOperations(operations []string) int {
    res:=0
    for _,val:=range operations{
        if val=="++X" || val== "X++"{
            res++
        }else{
            res--
        }
    }
    return res
}