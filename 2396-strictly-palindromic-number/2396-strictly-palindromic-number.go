func isStrictlyPalindromic(n int) bool {
    base:=2
    for base<=n-2{
        temp:=n
        str:=""
        for temp>0{
            str+=string(temp%base)
            temp=temp/base
        }
        if str!=reverseString(str){
            return false 
        }
        base++
    }
    return true 
}




func reverseString(s string) string {
    runes := []rune(s)
    i, j := 0, len(runes)-1

    for i < j {
        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }

    return string(runes)
}