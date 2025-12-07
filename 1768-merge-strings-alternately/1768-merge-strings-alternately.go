func mergeAlternately(word1 string, word2 string) string {
    i:=0
    j:=0
    res:=""
    iter:=0
    for iter<len(word1)+len(word2){
        if i==j && i<len(word1){
            res+=string(word1[i])
            i++
            if j>=len(word2){
                j++
            }
        }else if j<len(word2){
            res+=string(word2[j])
            j++
            if i>=len(word1){
                i++
            }
        }
        iter++
    }
    return res 
}