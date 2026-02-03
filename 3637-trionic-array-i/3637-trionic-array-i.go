func isTrionic(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }

    i := 0

    for i < n-1 && nums[i] < nums[i+1] {
        i++
    }
    if i == 0 || i == n-1 {
        return false
    }

    j := i
    for j < n-1 && nums[j] > nums[j+1] {
        j++
    }
    if j == i || j == n-1 {
        return false
    }

    for j < n-1 {
        if nums[j] >= nums[j+1] {
            return false
        }
        j++
    }

    return true
}
