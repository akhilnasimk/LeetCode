func recoverOrder(order []int, friends []int) []int {
    fMap:=make(map[int]bool)
    for _,val :=range friends {
        fMap[val]=true
    }
    res := []int {}
    // fmt.Println(fMap)
    for _,val:=range order{
        if fMap[val]{
            res=append(res,val)
        }
    }
    return res
}