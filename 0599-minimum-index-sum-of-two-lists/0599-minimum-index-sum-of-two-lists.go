func findRestaurant(list1 []string, list2 []string) []string {
   common :=make(map[string]int)
   L1Map:=make(map[string]int) 
   for ind,val:=range list1{
        L1Map[val]=ind
   }
    leastsum:=10000
   for ind,val:=range list2{
    if indexi,exist:=L1Map[val];exist{
        common[val]=indexi+ind
        if indexi+ind<leastsum{
            leastsum=indexi+ind
        }
    }
   }
   res:=[]string{}
   for key,val:=range common{
    if val==leastsum{
        res=append(res,key)
    }
   }
   return res
}