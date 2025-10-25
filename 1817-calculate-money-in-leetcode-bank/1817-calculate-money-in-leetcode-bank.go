func totalMoney(n int) int {
    i:=0
    start:=1
    total:=0
    for i < n {
        temp:=start 
        j:=0
        for j<7 && i<n{
            fmt.Println(temp)
            total+=temp
            temp++
            j++
            i++
        }
        start++
    }
    return total  
}