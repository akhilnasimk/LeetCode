func transformArray(nums []int) []int {
    res:=[]int {}
    for _,val:=range nums{
        res=append(res,val%2)
    }
    k:=0
    for ind,val:=range res{
        if val==0{
            res[ind],res[k]=res[k],res[ind]
            k++
        }
    }
    fmt.Println(res)
    return res
}