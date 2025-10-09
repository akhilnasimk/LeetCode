func singleNumber(nums []int) []int {
    count:= make(map[int]int)
    for _,val :=range nums{
        count[val]++
    }
    // fmt.Println(count)
    res := []int{}
    for key,val:=range count{
        if val==1{
            res=append(res,key)
        }
    }
    return res
}