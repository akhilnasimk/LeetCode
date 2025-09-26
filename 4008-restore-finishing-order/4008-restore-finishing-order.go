func recoverOrder(order []int, friends []int) []int {
    fMap:=make(map[int]int)
    for _,val :=range friends {
        fMap[val]=1
    }
    res := []int {}
    // fmt.Println(fMap)
    for _,val:=range order{
        if fMap[val]!=0{
            res=append(res,val)
        }
    }
    return res
}