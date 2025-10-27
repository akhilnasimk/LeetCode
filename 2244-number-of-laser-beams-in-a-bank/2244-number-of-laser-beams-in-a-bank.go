func numberOfBeams(bank []string) int {
    total:=0
    count:=0
    for _,val:=range bank {
        count1:=0
        for _,v:=range val {
            if string(v)=="1"{
                count1++
            }
        }
        if count1==0{
            continue
        }else if count==0{
            count=count1 
        }else{
            total+=count1*count
            count=count1
        }
    }
    return total 
}