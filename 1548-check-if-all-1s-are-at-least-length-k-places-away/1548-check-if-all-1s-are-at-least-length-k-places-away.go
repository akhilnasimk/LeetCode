func kLengthApart(nums []int, k int) bool {
    count:=0
    flag:=true 
    for _,val:=range nums {
        if val==0{
            count++
        }else{
            fmt.Println(count)
            if count<k && !flag{
                
                return false
            }
            count=0
            flag=false 
        }
    }
    return true 
}