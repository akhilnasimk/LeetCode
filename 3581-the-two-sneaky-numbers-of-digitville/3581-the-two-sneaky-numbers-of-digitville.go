func getSneakyNumbers(nums []int) []int {
    store:=make(map[int]int)
    res :=[]int{}
    for _,val :=range nums{
        store[val]++
        if store[val]==2{
            res=append(res,val)
        }
    }
    return res 

}