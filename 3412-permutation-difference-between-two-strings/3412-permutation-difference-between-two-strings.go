func findPermutationDifference(s string, t string) int {
    sMap:=make(map[string]int,len(s))
    tMap:=make(map[string]int,len(t))
    res:=0.0
    for i :=range len(s){
        sMap[string(s[i])]=i
        tMap[string(t[i])]=i
    }
    for key,_ :=range tMap{
        res+=(math.Abs(float64(sMap[key]-tMap[key])))
    }
    return int(res)
}