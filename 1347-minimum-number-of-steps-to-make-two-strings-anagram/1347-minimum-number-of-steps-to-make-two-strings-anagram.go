func minSteps(s string, t string) int {
    map1:=make(map[string]int)
    for _,val:=range t {
        map1[string(val)]++
    }
    count:=0
    for _,val:=range s {
        v,exist:=map1[string(val)]
        if exist && v>0{
            count++
            map1[string(val)]--
        }
    }
    return len(s)-count
}