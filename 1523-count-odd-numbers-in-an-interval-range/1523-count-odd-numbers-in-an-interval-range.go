func countOdds(low int, high int) int {
    count:=0
    var i int 
    if low%2!=0{
        i=low 
    }else{
        i=low+1
    }
    for i<=high{
        count++
        i+=2
    }
    return count 
}