func flipAndInvertImage(image [][]int) [][]int {
    for _,val:=range image{
        start:=0
        end:=len(val)-1
        for start<=end {
            val[start],val[end]=Convert(val[end]),Convert(val[start])
            start++
            end--
        }
    }
    return image
}

func Convert(val int)int{
    if val==0{
        return 1 
    }
    return 0 
}