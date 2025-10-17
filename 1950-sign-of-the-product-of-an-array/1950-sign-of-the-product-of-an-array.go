func arraySign(nums []int) int {
    sign := 1
    for _, n := range nums {
        if n == 0 {
            return 0
        }
        if n < 0 {
            sign = -sign
        }
    }
    return sign

}

