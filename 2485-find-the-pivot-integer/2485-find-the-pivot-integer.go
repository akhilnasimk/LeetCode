func pivotInteger(n int) int {
	sum := (n * (n + 1)) / 2

	curSum := 0

	for i := 1; i <= n; i++ {

		curSum += i

		rightSum := sum - curSum + i

		if curSum == rightSum {
			return i
		}
	}

	return -1
}