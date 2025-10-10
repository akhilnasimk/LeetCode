func canAliceWin(nums []int) bool {
    digi1:=0
    digi2:=0
    for _,val:=range nums{
        if val>9{
            digi2+=val
        }else{
            digi1+=val
        }
    }
    if digi1>digi2 || digi1<digi2{
        return true 
    }
    return false 
}