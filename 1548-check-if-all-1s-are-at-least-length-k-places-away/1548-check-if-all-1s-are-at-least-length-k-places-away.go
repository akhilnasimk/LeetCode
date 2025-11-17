func kLengthApart(nums []int, k int) bool {
	last := -1
	for i, val := range nums {
		if val == 1 {
			if (i-last)-1 < k && last != -1 {
				return false
			}
			last = i
		}
	}
	return true
}