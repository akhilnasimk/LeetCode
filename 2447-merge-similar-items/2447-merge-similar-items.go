func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    item2Map := make(map [int]int)
    for _,val:= range items2{
        item2Map[val[0]]= val[1]
    }
    for _,val:= range items1{
        if res,ok:=item2Map[val[0]];ok{
            item2Map[val[0]]=res+val[1]
        }else{
            item2Map[val[0]]=val[1]
        }
    }
    fmt.Println(item2Map)
    resl := make([][]int,0)
    for id,val:=range item2Map{
        resl=append(resl,[]int{id,val})
    }
    sort.Slice(resl, func(i, j int) bool {
        return resl[i][0] < resl[j][0]
    })
    return resl
}