func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    count:=0
    for _,val:=range nums1{
        for _,v:=range nums2{
            if val%(v*k)==0{
                    count++
            }
        }
    }   
    return count 
    }
    