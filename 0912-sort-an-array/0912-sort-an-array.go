func sortArray(nums []int) []int {
    if len(nums)==1{
        return nums
    }

    mid:=len(nums)/2
    left:=sortArray(nums[:mid])
    right:=sortArray(nums[mid:])

    return merge(left,right)
}

func merge(left,right []int)[]int{
    merged:=[]int{}
    i:=0
    j:=0
    for i<len(left)&&j<len(right){
        if left[i]<right[j]{
            merged=append(merged,left[i])
            i++
        }else{
            merged=append(merged,right[j])
            j++
        }
    }
    if i<len(left){
       merged = append(merged, left[i:]...)
    }
    if j<len(right){
        merged = append(merged, right[j:]...)
    }

    return merged 
}