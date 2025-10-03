func findPermutationDifference(s string, t string) int {
    sMap:=make(map[string]int,len(s))
    tMap:=make(map[string]int,len(t))
    res:=0.0
    for ind,val:=range s{
        sMap[string(val)]=ind
    }
    for ind,val:=range t{
        tMap[string(val)]=ind
    }
    for key,_ :=range tMap{
        res+=(math.Abs(float64(sMap[key]-tMap[key])))
    }
    return int(res)
}