func threeConsecutiveOdds(arr []int) bool {
    first:=0
    second:=1 
    thired:=2
    flag:=false 
    for thired < len(arr){
        if arr[first]%2!=0 &&arr[second]%2!=0 &&arr[thired]%2!=0{
            flag= true 
            break 
        }
        first++
        second++
        thired++
    }
    if flag {
        return true 
    }
    return false 
}