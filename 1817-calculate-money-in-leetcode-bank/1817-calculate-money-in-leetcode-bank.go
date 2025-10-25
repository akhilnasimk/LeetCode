func totalMoney(n int) int {
    start:=1
    each:=start 
    i:=0
    total :=0
    for i<n{
        if each>start+6{
            start++
            each=start 
        }else{
            fmt.Println(each)
            total+=each 
            each++
            i++
        }
    } 
    return total 
}