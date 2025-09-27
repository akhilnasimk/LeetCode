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
    for i:=temp+1;i<len(word);i++{
        res=append(res,word[i])
    }

    fmt.Println(word)
    return string(res) 
}