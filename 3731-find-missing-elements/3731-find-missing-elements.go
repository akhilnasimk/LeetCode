func findMissingElements(nums []int) []int {
    missing:=[]int{}
    elements:=make(map[int]bool)
    low:=101
    heigh:=0
    for _,val:=range nums{
        elements[val]=true 
        if val<low{
            low=val
        }
        if heigh <val{
            heigh=val
        }
    }

    for low<heigh {
        if !elements[low]{
            missing=append(missing,low)
        }
        low++
    }
    fmt.Println(low,heigh)
    return missing 

}