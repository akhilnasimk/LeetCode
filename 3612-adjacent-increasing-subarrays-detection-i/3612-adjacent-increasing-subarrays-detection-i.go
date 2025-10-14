func hasIncreasingSubarrays(nums []int, k int) bool {
    if k==1{
        return true 
    }
    i:=0
    flag:=false 
    for i<=len(nums)-k{
        strict:=true    
        for j:=i;j<i+k-1;j++{
            fmt.Println(nums[j],nums[j+1])
            if  nums[j] >= nums[j+1]{
                strict=false   
                break 
            }
        }
        if flag==true && strict==true {
            return true 
        }else if flag==true && strict==false{
            flag=false  
            i=i-k+1
        }else if strict {
            fmt.Println(i)
            flag=true 
            i+=k
        }else{
            i++
        }
        fmt.Println(flag)

    }
    return false 
}