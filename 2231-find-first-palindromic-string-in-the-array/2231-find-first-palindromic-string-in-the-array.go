func firstPalindrome(words []string) string {
    var res string 
    for _,val:=range words{
        left:=0 
        right:=len(val)-1
        flag:=true   
        for right>=left {
            if val[left]==val[right]{
                flag=flag 
                left++
                right--
            }else{
                flag=false 
                left++
                right--
            }
        }
        if flag{
            res=val
            break 
        }
    }
    return res
}