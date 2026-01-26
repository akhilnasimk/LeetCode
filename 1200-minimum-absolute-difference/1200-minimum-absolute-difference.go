func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

    mindiff:=100001
    res:=[][]int{}

    for i:=0;i<len(arr)-1;i++{
        absdif:=Abs(arr[i]-arr[i+1])
        if absdif<mindiff{
            mindiff=absdif
            res=[][]int{[]int{arr[i],arr[i+1]}}
            continue 
        }
        if absdif==mindiff{
            res=append(res,[]int{arr[i],arr[i+1]})
        }
    }
    return res
}
func Abs(a int )int{
    if a <0 {
        return -a 
    }
    return a 
}