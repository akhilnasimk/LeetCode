func smallestRepunitDivByK(k int) int {
    if k%2 == 0 || k%5 == 0 {
        return -1
    }

    remainder := 1 % k
    length := 1

    for remainder != 0 {
        remainder = (remainder*10 + 1) % k
        length++
    }

    return length
}
