func removeAnagrams(words []string) []string {
    i:=1
    for i<len(words){
        res:=IsAnagram(words[i],words[i-1])
        if res{
            words=append(words[:i],words[i+1:]...)
            continue 
        }
        i++
    }
    return words 
}

func IsAnagram(s,old string  )bool{
    if len(s)!=len(old){
        return false
    }
        ol := []rune(old)
	    sort.Slice(ol, func(i, j int) bool { return ol[i] < ol[j] })
        r := []rune(s)
	    sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
        if string(ol)==string(r){
            return true 
        }
    return false 
}