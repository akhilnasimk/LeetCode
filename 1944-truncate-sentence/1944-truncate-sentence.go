func truncateSentence(s string, k int) string {
    count:=0 
    res:=""
    for _,val:=range s{
        if string(val)==" "{
            count++
        }
        if k==count{
            break 
        }
        res+=string(val)
    }
    return res
}