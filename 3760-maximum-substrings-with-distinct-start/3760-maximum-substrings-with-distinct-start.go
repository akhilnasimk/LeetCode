func maxDistinct(s string) int {
    exist:=make(map[rune]bool)
    count:=0
    for _,val:=range s {
        if yes:=exist[val];yes{
            continue 
        }else{
            count++
            exist[val]=true 
        }
    }

    return count 
}