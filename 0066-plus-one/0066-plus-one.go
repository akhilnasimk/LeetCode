func plusOne(digits []int) []int {
    addCarry(&digits, len(digits)-1)
    return digits
}

func addCarry(digits *[]int, i int) {
    if i < 0 {
        *digits = append([]int{1}, *digits...)
        return
    }

    if (*digits)[i] < 9 {
        (*digits)[i]++
        return
    }

    (*digits)[i] = 0
    addCarry(digits, i-1)
}
