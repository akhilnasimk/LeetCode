func findCenter(edges [][]int) int {
    uni:=make(map[int]int)
    for _,val :=range edges{
        for _,iner:=range val{
            uni[iner]++
        }
    }
    res:=0
    for key,val :=range uni{
        if val==len(edges){
            res=key
        }
    }
    return res
}