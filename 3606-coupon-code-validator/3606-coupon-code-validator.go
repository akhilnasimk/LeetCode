func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
    re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
    res:=[]string{}
    map1:=make(map[string][]string)
    for ind,val:=range code {
        if re.MatchString(val){
            if (businessLine[ind]=="electronics" || businessLine[ind]=="grocery" || businessLine[ind]=="pharmacy"  || businessLine[ind]=="restaurant" )&& isActive[ind]==true{
                fmt.Println(val,isActive[ind])
                map1[ businessLine[ind]] =append(map1[ businessLine[ind]],val)
            }
        }
    } 
    fmt.Println(map1)
    order:=[]string{"electronics", "grocery", "pharmacy", "restaurant"}  
    for _,val:=range order{
        sort.Strings(map1[val])
        res = append(res, map1[val]...)
    }
    return res
}