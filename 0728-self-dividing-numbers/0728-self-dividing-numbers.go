func selfDividingNumbers(left int, right int) []int {
    res:=[]int{}
    for i:=left;i<=right;i++{
        temp:=i
        flag:=true 
        for temp>0{
            rem:=temp%10
            if rem==0 || i%rem!=0{
                flag=false 
                break 
            }
            temp=temp/10
        }
        if flag{
            res=append(res,i)
        }
    }
    return res
}