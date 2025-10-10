

func countKeyChanges(s string) int {
    s=strings.ToLower(s)
    count:=0
    for i,val:=range s {
        if i<len(s)-1{
            if string(val)!= string(s[i+1]){
            count++
        }
        }
    }
    return count
}