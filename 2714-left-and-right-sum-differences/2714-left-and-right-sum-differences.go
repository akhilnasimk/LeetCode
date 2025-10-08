func leftRightDifference(nums []int) []int {
    leftsum:=[]int{}
    sum:=0
    i:=0
    leftsum=append(leftsum,sum)
    for i<len(nums)-1{
        sum+=nums[i]
        i++
        leftsum=append(leftsum,sum)
    }
    sum=0
    rightsum:=[]int{}
    j:=len(nums)-1
    rightsum=append(rightsum,sum)
    for j>0{
        sum+=nums[j]
        j--
        rightsum=append(rightsum,sum)
    }
    rp:=len(nums)-1
    res:=[]int{}
    for _,val:=range leftsum{
        res=append(res,AbsSub(val,rightsum[rp]))
        // fmt.Println(val,rightsum[rp])
        rp--
    }
    // fmt.Println(leftsum,rightsum)
    return res
}

func AbsSub(a,b int)int{
    n:=a-b
    if n<0{
        return -n 
    }
    return n 
}