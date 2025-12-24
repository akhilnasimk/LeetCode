func minimumBoxes(apple []int, capacity []int) int {
    sorted:=MergeSort(capacity)
    count:=0
    sum:=0
    for _,val:=range apple{
        sum+=val
    }
    i:=0
    for sum>0{
        sum-=sorted[i]
        i++
        count++
    }
    return count 
}


func MergeSort(ar []int)[]int{
    if len(ar)<=1{
        return ar 
    }

    middle:=len(ar)/2
    left:=MergeSort(ar[:middle])
    right:=MergeSort(ar[middle:])

    return Merge(left,right)
}

func Merge(left ,right []int)[]int{
    merged:=[]int{}
    i:=0
    j:=0
    for i<len(left) && j<len(right){
        if left[i]>right[j]{
            merged=append(merged,left[i])
            i++
        }else{
            merged=append(merged,right[j])
            j++
        }
    }

    merged = append(merged, left[i:]...)
    merged = append(merged, right[j:]...)
    return merged 
}