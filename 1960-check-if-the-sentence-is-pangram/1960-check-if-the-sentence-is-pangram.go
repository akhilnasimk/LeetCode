func checkIfPangram(sentence string) bool {
    alpha := make(map[rune]bool)
    for _,val:=range sentence {
        alpha[val]=true 
    }

    if len(alpha)==26 {
        return true 
    }
    return false 
}