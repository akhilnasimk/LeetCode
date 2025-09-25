func maxFreqSum(s string) int {
    vowel:=make(map[rune]int)
    cons:=make(map[rune]int)
    for _, val := range s {
        if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
            vowel[val]++
        } else {
            cons[val]++
        }
    }
    fmt.Println(vowel,"  ",cons)
    BV:=0
    for _,val :=range vowel{
        if val>BV{
            BV=val
        }
    }
    CV:=0
    for _,val :=range cons{
        if val>CV{
            CV=val
        }
    }
    return BV+CV
}