func hasSameDigits(s string) bool {
    if len(s)==0 ||  len(s)==1{
        return true 
    }
    arr:=[]int{}
    for _,val:=range s {
        newVAl,_:=strconv.Atoi(string(val))
        arr=append(arr,newVAl)
    }
    // fmt.Println(arr)
    for len(arr)>2 {
        for i:=0;i<len(arr)-1;i++{
            arr[i]=(arr[i]+arr[i+1])%10
        }
        arr=arr[:len(arr)-1]
    }
    if arr[0]==arr[1]{
        return true 
    }
    return false 
}