
func numTilePossibilities(tiles string) int {
    freq := make(map[rune]int)

    for _, ch := range tiles {
        freq[ch]++
    }

    var backtrack func() int
    backtrack = func() int {
        count := 0

        for ch := range freq {
            if freq[ch] > 0 {
                freq[ch]--


                count += 1


                count += backtrack()
                freq[ch]++
            }
        }

        return count
    }

    return backtrack()
}
