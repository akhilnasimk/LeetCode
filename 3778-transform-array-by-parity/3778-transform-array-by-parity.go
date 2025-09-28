func transformArray(nums []int) []int {
    k := 0 
    for i := 0; i < len(nums); i++ {
        mod := nums[i] % 2 
        nums[i] = mod     

        if mod == 0 { 
            nums[i], nums[k] = nums[k], nums[i]
            k++
        }
    }
    return nums
}
