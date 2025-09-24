func scoreOfString(s string) int {
    var sum int 
    abs:=func (n int )(int) {
        if n<0{
            return -n 
        }
        return n
    }
    for i:=0 ;i<len(s)-1;i++{
            sum+=abs(int(s[i])-int(s[i+1]))
    }
    return sum
}