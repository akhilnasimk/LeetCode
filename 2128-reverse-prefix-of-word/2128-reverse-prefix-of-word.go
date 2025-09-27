func reversePrefix(word string, ch byte) string {
    left:=0
    res:=[]byte {}
    for ind,val:=range word{
        if byte(val)==ch{
            left=ind
            break
        }
    }
    fmt.Println(left)
    temp:=left
    for 0<=left{
        res=append(res,word[left])
        left--
    }
    res=append(res,word[temp+1:len(word)]...)

    fmt.Println(word)
    return string(res) 
}